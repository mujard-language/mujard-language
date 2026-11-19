package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// samdSteps builds the family:samd recipe instance. The classic Arduino SAMD21
// core is the simplest ARM family here: compile cores/arduino (with its api/USB
// subdirs) + the variant's sources, link everything + core.a inside one
// --start-group against the CMSIS cortexM math lib, objcopy to .bin/.hex. Flags
// match the arduino-cli verbose recipe for arduino:samd:arduino_zero_native
// (core 1.8.14, arm-none-eabi gcc 7-2017q4, CMSIS 4.5.0). Board-specific values
// arrive via manifest vars.
func (p *Plan) samdSteps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	v := p.Vars
	for _, k := range []string{
		"mcu", "arch", "boardDefine", "chipDefine", "fCpu", "ideVersion", "cpuFlag",
		"usbVid", "usbPid", "usbProduct", "usbManufacturer", "ldscript", "mathLib", "mathLibDir",
	} {
		if v[k] == "" {
			return fmt.Errorf("%s: samd manifest needs var %q", b.Short, k)
		}
	}

	gccBin := pkg.Tools["gccBin"]
	cmsis := pkg.Tools["cmsis"]
	cmsisAtmel := pkg.Tools["cmsisAtmel"]
	if gccBin == "" || cmsis == "" || cmsisAtmel == "" {
		return &unsupported{fmt.Sprintf("package %s needs gccBin/cmsis/cmsisAtmel tool dirs", pkg.ID)}
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
		"-DF_CPU=" + v["fCpu"],
		"-DARDUINO=" + v["ideVersion"],
		"-DARDUINO_" + v["boardDefine"],
		"-DARDUINO_ARCH_" + v["arch"],
		"-D" + v["chipDefine"],
		"-DUSB_VID=" + v["usbVid"],
		"-DUSB_PID=" + v["usbPid"],
		"-DUSBCON",
		`-DUSB_MANUFACTURER="` + v["usbManufacturer"] + `"`,
		`-DUSB_PRODUCT="` + v["usbProduct"] + `"`,
	}
	for _, d := range b.Defines {
		defs = append(defs, "-D"+d)
	}
	inc := []string{
		"-I" + j.SketchDir,
		"-I" + filepath.Join(cmsis, "Include"),
		"-I" + cmsisAtmel,
		"-I" + filepath.Join(coreDir, "api/deprecated"),
		"-I" + filepath.Join(coreDir, "api/deprecated-avr-comp"),
		"-I" + coreDir,
		"-I" + variantDir,
	}
	for _, li := range libIncs {
		inc = append(inc, "-I"+li)
	}
	base := []string{v["cpuFlag"], "-mthumb", "-c", "-g", "-Os", "-w",
		"-ffunction-sections", "-fdata-sections", "-nostdlib",
		"--param", "max-inline-insns-single=500", "-MMD"}
	cflags := append(append(append([]string{}, base...), "-std=gnu11"), defs...)
	cxxflags := append(append(append([]string{}, base...), "-std=gnu++11",
		"-fno-threadsafe-statics", "-fno-rtti", "-fno-exceptions"), defs...)
	sflags := append(append(append([]string{}, v["cpuFlag"], "-mthumb", "-c",
		"-x", "assembler-with-cpp", "-w", "-MMD"), defs...), inc...)
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

	elf := filepath.Join(j.OutDir, "firmware.elf")
	link := []string{
		"-L" + j.OutDir,
		"-Os", "-Wl,--gc-sections",
		"-T" + pkg.subRoot(v["ldscript"]),
		"-Wl,-Map," + filepath.Join(j.OutDir, "firmware.map"),
		"--specs=nano.specs", "--specs=nosys.specs",
		v["cpuFlag"], "-mthumb",
		"-Wl,--cref", "-Wl,--check-sections", "-Wl,--unresolved-symbols=report-all",
		"-Wl,--warn-common", "-Wl,--warn-section-align",
		"-o", elf,
	}
	link = append(link, sketchObjs...)
	link = append(link,
		"-Wl,--start-group",
		"-L"+filepath.Join(cmsis, v["mathLibDir"]),
		"-l"+v["mathLib"], "-lm",
		coreA,
		"-Wl,--end-group")
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cxx, Args: link})

	bin := filepath.Join(j.OutDir, "firmware.bin")
	hex := filepath.Join(j.OutDir, "firmware.hex")
	p.Steps = append(p.Steps, Step{Name: "bin", Exe: objcopy, Args: []string{"-O", "binary", elf, bin}})
	p.Steps = append(p.Steps, Step{Name: "hex", Exe: objcopy, Args: []string{"-O", "ihex", "-R", ".eeprom", elf, hex}})
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}
