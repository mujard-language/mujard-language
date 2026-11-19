package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// renesasSteps builds the family:renesas recipe instance. The ArduinoCore-renesas
// (UNO R4) recipe: core + variant sources compile against the variant's own
// response files (defines.txt / includes.txt / cxxflags.txt live in the variant
// dir and carry the whole FSP include chain via -iprefix<root>), the prebuilt
// libfsp.a ships in variant/libs, and the link wraps it + core.a in
// --whole-archive/--start-group. objcopy takes only .text/.data for the .bin.
// Flags match the arduino-cli verbose recipe for arduino:renesas_uno:minima
// (core 1.6.0, arm-none-eabi gcc 7-2017q4). Board-specific values arrive via
// manifest vars.
func (p *Plan) renesasSteps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	v := p.Vars
	for _, k := range []string{
		"mcu", "arch", "boardDefine", "variantDefine", "fCpu", "ideVersion",
		"cpuFlag", "floatFlags", "tinyusbMcu", "definesFile", "includesFile",
		"ldscript", "fspLib",
	} {
		if v[k] == "" {
			return fmt.Errorf("%s: renesas manifest needs var %q", b.Short, k)
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
	fspLib := filepath.Join(variantDir, v["fspLib"])
	if st, err := os.Stat(fspLib); err != nil || st.IsDir() {
		return &unsupported{fmt.Sprintf("prebuilt libfsp missing at %s", fspLib)}
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
		"-DF_CPU=" + v["fCpu"],
		"-D" + v["boardDefine"],
		"-DARDUINO=" + v["ideVersion"],
		`-DPROJECT_NAME="` + filepath.Join(j.OutDir, "firmware") + `"`,
		"-D" + v["variantDefine"],
		"-DARDUINO_ARCH_" + v["arch"],
		"-DARDUINO_ARCH_RENESAS",
		"-DARDUINO_FSP",
		"-D_XOPEN_SOURCE=700",
		"@" + filepath.Join(variantDir, v["definesFile"]),
		"-DCFG_TUSB_MCU=" + v["tinyusbMcu"],
	}
	for _, d := range b.Defines {
		defs = append(defs, "-D"+d)
	}
	inc := []string{
		"-I" + j.SketchDir,
		"-I" + filepath.Join(coreDir, "tinyusb"),
		"-I" + filepath.Join(coreDir, "api/deprecated"),
		"-I" + filepath.Join(coreDir, "api/deprecated-avr-comp"),
		"-I" + coreDir,
		"-I" + variantDir,
		"-iprefix" + pkg.Root,
		"@" + filepath.Join(variantDir, v["includesFile"]),
	}
	for _, li := range libIncs {
		inc = append(inc, "-I"+li)
	}
	base := []string{"-c", "-w", "-Os", "-g3", "-nostdlib", "-MMD",
		"-fsigned-char", "-ffunction-sections", "-fdata-sections",
		"-fmessage-length=0", "-fno-builtin"}
	for _, f := range strings.Fields(v["floatFlags"]) {
		base = append(base, f)
	}
	base = append(base, v["cpuFlag"], "-mthumb")
	cflags := append(append(append([]string{}, base...), "-std=gnu11"), defs...)
	cxxflags := append(append(append([]string{}, base...), "-std=gnu++17",
		"-fno-use-cxa-atexit", "-fno-rtti", "-fno-exceptions"), defs...)
	sflags := append(append(append([]string{}, "-c", "-w", "-x", "assembler-with-cpp",
		"-MMD", v["cpuFlag"], "-mthumb"), defs...), inc...)
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
	var floatFlags []string
	for _, f := range strings.Fields(v["floatFlags"]) {
		floatFlags = append(floatFlags, f)
	}
	link := append([]string{
		"-Wl,--gc-sections", "--specs=nosys.specs", "-w",
		v["cpuFlag"], "-mthumb",
	}, floatFlags...)
	link = append(link,
		"-o", elf,
		"-L"+j.OutDir,
		"-L"+variantDir,
		"-T"+filepath.Join(variantDir, v["ldscript"]),
	)
	link = append(link, sketchObjs...)
	link = append(link,
		"-Wl,--whole-archive", "-Wl,--start-group",
		fspLib, coreA,
		"-Wl,--no-whole-archive",
		"--specs=nano.specs", "-lstdc++", "-lsupc++", "-lm", "-lc", "-lgcc", "-lnosys",
		"-Wl,--end-group",
		"-Wl,-Map,"+filepath.Join(j.OutDir, "firmware.map"))
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cxx, Args: link})

	bin := filepath.Join(j.OutDir, "firmware.bin")
	hex := filepath.Join(j.OutDir, "firmware.hex")
	p.Steps = append(p.Steps, Step{Name: "bin", Exe: objcopy, Args: []string{"-O", "binary", "-j", ".text", "-j", ".data", elf, bin}})
	p.Steps = append(p.Steps, Step{Name: "hex", Exe: objcopy, Args: []string{"-O", "ihex", "-j", ".text", "-j", ".data", elf, hex}})
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}
