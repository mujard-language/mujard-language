package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// mbedSteps builds the family:mbed recipe instance. The ArduinoCore-mbed
// (Nano 33 BLE) recipe: core + variant.cpp compile against the variant's own
// response files (defines.txt / cflags.txt / cxxflags.txt / includes.txt via
// -iprefix<cores/arduino>), the linker script is gcc-preprocessed
// (-E -P -x c) into the build dir, and the link wraps core.a + the variant's
// prebuilt mbed/cc310 archives in --whole-archive. objcopy makes .bin + .hex.
// Flags match the arduino-cli verbose recipe for arduino:mbed_nano:nano33ble
// (core 4.6.0, arm-none-eabi gcc 7-2017q4). Board-specific values arrive via
// manifest vars.
func (p *Plan) mbedSteps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	v := p.Vars
	for _, k := range []string{
		"mcu", "arch", "boardDefine", "archChipDefine", "fCpu", "ideVersion",
		"cpuFlag", "floatFlags", "definesFile", "includesFile", "ldflagsFile",
		"ldscriptIn", "variantLibs",
	} {
		if v[k] == "" {
			return fmt.Errorf("%s: mbed manifest needs var %q", b.Short, k)
		}
	}

	gccBin := pkg.Tools["gccBin"]
	if gccBin == "" {
		return &unsupported{fmt.Sprintf("package %s has no gccBin tool dir", pkg.ID)}
	}
	exeAt := func(name string) (string, bool) {
		for _, c := range []string{filepath.Join(gccBin, name), filepath.Join(gccBin, name+".exe")} {
			if st, err := os.Stat(c); err == nil && !st.IsDir() {
				return c, true
			}
		}
		return "", false
	}
	cc, ok1 := exeAt("arm-none-eabi-gcc")
	cxx, ok2 := exeAt("arm-none-eabi-g++")
	ar, ok3 := exeAt("arm-none-eabi-ar")
	size, ok4 := exeAt("arm-none-eabi-size")
	objcopy, ok5 := exeAt("arm-none-eabi-objcopy")
	if !(ok1 && ok2 && ok3 && ok4 && ok5) {
		return &unsupported{fmt.Sprintf("arm-none-eabi toolchain missing under %s", gccBin)}
	}

	coreDir := pkg.subRoot(b.Core)
	variantDir := pkg.subRoot(b.Variant)
	var variantLibs []string
	for _, rel := range strings.Fields(v["variantLibs"]) {
		lp := filepath.Join(variantDir, rel)
		if st, err := os.Stat(lp); err != nil || st.IsDir() {
			return &unsupported{fmt.Sprintf("prebuilt variant lib missing at %s", lp)}
		}
		variantLibs = append(variantLibs, lp)
	}

	coreSrcs, err := sourcesRecursive(coreDir)
	if err != nil {
		return fmt.Errorf("core sources: %w", err)
	}
	varSrcs, err := sourcesRecursive(variantDir)
	if err == nil {
		coreSrcs = append(coreSrcs, varSrcs...)
	}
	sketchSrcs, inos, err := sources(j.SketchDir)
	if err != nil {
		return fmt.Errorf("sketch sources: %w", err)
	}
	gen, err := p.preprocessIno(inos)
	if err != nil {
		return fmt.Errorf("ino preprocessing: %w", err)
	}
	sketchSrcs = append(sketchSrcs, gen...)
	if len(sketchSrcs) == 0 {
		return fmt.Errorf("sketchDir %s has no .ino/.c/.cpp/.S sources", j.SketchDir)
	}
	p.CoreSrcs, p.SketchSrcs = coreSrcs, sketchSrcs

	libIncs, libSrcs, err := p.resolveLibs()
	if err != nil {
		return fmt.Errorf("libraries: %w", err)
	}
	if err := os.MkdirAll(j.OutDir, 0o755); err != nil {
		return err
	}

	defs := []string{
		"-D" + v["archChipDefine"],
		"-MMD",
		v["cpuFlag"], "-mfloat-abi=softfp", "-mfpu=fpv4-sp-d16",
		"-DARDUINO=" + v["ideVersion"],
		"-DARDUINO_" + v["boardDefine"],
		"-DARDUINO_ARCH_" + v["arch"],
		"-DARDUINO_ARCH_MBED",
		"-DARDUINO_LIBRARY_DISCOVERY_PHASE=0",
	}
	for _, d := range b.Defines {
		defs = append(defs, "-D"+d)
	}
	inc := []string{
		"-I" + j.SketchDir,
		"-I" + coreDir,
		"-I" + variantDir,
		"-I" + filepath.Join(coreDir, "api/deprecated"),
		"-I" + filepath.Join(coreDir, "api/deprecated-avr-comp"),
		"-iprefix" + coreDir,
		"@" + filepath.Join(variantDir, v["includesFile"]),
	}
	for _, li := range libIncs {
		inc = append(inc, "-I"+li)
	}
	base := []string{"-c", "-w", "-g3", "-nostdlib",
		"@" + filepath.Join(variantDir, v["definesFile"])}
	cflags := append(append(append([]string{}, base...),
		"@"+filepath.Join(variantDir, "cflags.txt")), defs...)
	cxxflags := append(append(append([]string{}, base...),
		"@"+filepath.Join(variantDir, "cxxflags.txt")), defs...)
	sflags := append(append(append([]string{}, "-c", "-w", "-x", "assembler-with-cpp",
		"-MMD", v["cpuFlag"]), defs...), inc...)
	cflags = append(cflags, inc...)
	cxxflags = append(cxxflags, inc...)

	objRoot := filepath.Join(j.OutDir, "obj")
	coreA := filepath.Join(j.OutDir, "core.a")
	coreObjRoot := filepath.Join(objRoot, "core")
	warm := false
	if !j.NoCache {
		if final := p.coreCacheDir(cflags, cxxflags, sflags); final != "" {
			if _, err := os.Stat(filepath.Join(final, "core.a")); err == nil {
				coreA, warm = filepath.Join(final, "core.a"), true
			} else {
				p.CoreFinal, p.CoreStaging = final, p.coreStaging(final)
				coreObjRoot = filepath.Join(p.CoreStaging, "obj")
				coreA = filepath.Join(p.CoreStaging, "core.a")
			}
		}
	}

	objs := objNamer{}
	add := func(root, src string) string {
		obj := objs.path(root, src)
		var args []string
		switch strings.ToLower(filepath.Ext(src)) {
		case ".c":
			args = append(args, cflags...)
		case ".s":
			args = append(args, sflags...)
		default:
			args = append(args, cxxflags...)
		}
		args = append(args, src, "-o", obj+".part")
		p.Steps = append(p.Steps, Step{Name: "compile " + filepath.Base(src), Exe: compilerFor(src, cc, cxx), Parallel: true, Out: obj, Args: args})
		return obj
	}

	var coreObjs []string
	if !warm {
		for _, src := range coreSrcs {
			coreObjs = append(coreObjs, add(coreObjRoot, src))
		}
		p.Steps = append(p.Steps, Step{Name: "archive core", Exe: ar, Args: append([]string{"cr", coreA}, coreObjs...)})
	}
	var sketchObjs []string
	for _, src := range sketchSrcs {
		sketchObjs = append(sketchObjs, add(filepath.Join(objRoot, "sketch"), src))
	}
	for name, srcs := range libSrcs {
		for _, src := range srcs {
			sketchObjs = append(sketchObjs, add(filepath.Join(objRoot, "libs", name), src))
		}
	}

	// The linker script is gcc-preprocessed so its #define'd memory layout resolves.
	ldOut := filepath.Join(j.OutDir, "linker_script.ld")
	p.Steps = append(p.Steps, Step{Name: "linker script", Exe: cxx,
		Args: []string{"-E", "-P", "-x", "c", filepath.Join(variantDir, v["ldscriptIn"]), "-o", ldOut}})

	elf := filepath.Join(j.OutDir, "firmware.elf")
	link := []string{
		"-L" + j.OutDir,
		"-Wl,--gc-sections", "-w", "-Wl,--as-needed",
		"@" + filepath.Join(variantDir, v["ldflagsFile"]),
		"-T" + ldOut,
		"-Wl,-Map," + filepath.Join(j.OutDir, "firmware.map"),
		"--specs=nosys.specs",
		"-o", elf,
	}
	link = append(link, sketchObjs...)
	link = append(link, "-Wl,--whole-archive", coreA)
	link = append(link, variantLibs...)
	link = append(link,
		"-Wl,--no-whole-archive",
		"-Wl,--start-group", "-lstdc++", "-lsupc++", "-lm", "-lc", "-lgcc", "-lnosys",
		"-Wl,--end-group")
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cxx, Args: link})

	bin := filepath.Join(j.OutDir, "firmware.bin")
	hex := filepath.Join(j.OutDir, "firmware.hex")
	p.Steps = append(p.Steps, Step{Name: "bin", Exe: objcopy, Args: []string{"-O", "binary", elf, bin}})
	p.Steps = append(p.Steps, Step{Name: "hex", Exe: objcopy, Args: []string{"-O", "ihex", "-R", ".eeprom", elf, hex}})
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}
