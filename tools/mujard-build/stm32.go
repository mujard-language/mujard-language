package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// stm32Steps builds the family:stm32 recipe instance. The STM32 Arduino core
// (stm32duino) compiles every source under cores/arduino (+ avr/stm32/api subdirs),
// the board's variant directory (its variant_*/PeripheralPins_* files and startup
// assembly), and the whole SrcWrapper HAL/LL library tree — the linker then trims
// with --gc-sections. Flags match the arduino-cli verbose recipe captured for
// STMicroelectronics:stm32:GenF1:pnum=BLUEPILL_F103C8 (core 3.0.0, xpack gcc 14.2.1).
// Board-specific values arrive via manifest vars.
func (p *Plan) stm32Steps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	v := p.Vars
	for _, k := range []string{
		"mcu", "arch", "boardDefine", "familyDefine", "chipDefine", "variantH",
		"uartModule", "fCpu", "ideVersion", "cpuFlag", "vectTab", "halDrivers",
		"familyInc", "halInc", "halSrc", "sysFam", "deviceSrc", "deviceSrcGcc",
		"usbDevInc", "usbDevSrc", "virtioInc", "openampInc", "openampInc2",
		"libmetalInc", "virtDriverInc", "ldscript", "ldscriptSys", "flashOffset",
	} {
		if v[k] == "" {
			return fmt.Errorf("%s: stm32 manifest needs var %q", b.Short, k)
		}
	}

	gccBin := pkg.Tools["gccBin"]
	cmsis := pkg.Tools["cmsis"]
	cmsisDsp := pkg.Tools["cmsisDsp"]
	if gccBin == "" || cmsis == "" || cmsisDsp == "" {
		return &unsupported{fmt.Sprintf("package %s needs gccBin/cmsis/cmsisDsp tool dirs", pkg.ID)}
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
	ar, ok3 := exeAt("arm-none-eabi-gcc-ar")
	size, ok4 := exeAt("arm-none-eabi-size")
	objcopy, ok5 := exeAt("arm-none-eabi-objcopy")
	if !(ok1 && ok2 && ok3 && ok4 && ok5) {
		return &unsupported{fmt.Sprintf("arm-none-eabi toolchain missing under %s", gccBin)}
	}

	coreDir := pkg.subRoot(b.Core)
	variantDir := pkg.subRoot(b.Variant)
	srcWrap := pkg.subRoot("libraries/SrcWrapper/src")
	srcWrapInc := pkg.subRoot("libraries/SrcWrapper/inc")

	coreSrcs, err := sourcesRecursive(coreDir)
	if err != nil {
		return fmt.Errorf("core sources: %w", err)
	}
	varSrcs, err := sourcesRecursive(variantDir)
	if err == nil {
		coreSrcs = append(coreSrcs, varSrcs...)
	}
	halSrcs, err := sourcesRecursive(srcWrap)
	if err != nil {
		return fmt.Errorf("SrcWrapper sources: %w", err)
	}
	coreSrcs = append(coreSrcs, halSrcs...)

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
	// arduino-cli writes requiredLibraries.cpp (#include <SrcWrapper.h>) so the
	// library solver pulls the HAL in; the direct backend compiles SrcWrapper
	// unconditionally, but still writes the file to keep the unit list faithful.
	if err := os.MkdirAll(j.OutDir, 0o755); err != nil {
		return err
	}
	reqLibs := filepath.Join(j.OutDir, "requiredLibraries.cpp")
	if err := os.WriteFile(reqLibs, []byte("#include <SrcWrapper.h>\n"), 0o644); err != nil {
		return err
	}
	sketchSrcs = append(sketchSrcs, reqLibs)
	p.CoreSrcs, p.SketchSrcs = coreSrcs, sketchSrcs

	libIncs, libSrcs, err := p.resolveLibs()
	if err != nil {
		return fmt.Errorf("libraries: %w", err)
	}

	// prebuild.sh in the reference recipe only writes build.opt holding one
	// -fmacro-prefix-map line; the backend writes the equivalent directly.
	buildOpt := filepath.Join(j.OutDir, "build.opt")
	if err := os.WriteFile(buildOpt, []byte(`-fmacro-prefix-map="`+pkg.Root+`"=.`+"\n"), 0o644); err != nil {
		return err
	}

	defs := []string{
		"-D" + v["familyDefine"],
		"-DARDUINO=" + v["ideVersion"],
		"-DARDUINO_GENERIC",
		"-DARDUINO_" + v["boardDefine"],
		"-DARDUINO_ARCH_" + v["arch"],
		`-DBOARD_NAME="` + v["boardDefine"] + `"`,
		`-DVARIANT_H="` + v["variantH"] + `"`,
		"-D" + v["chipDefine"],
		"-D" + v["uartModule"],
	}
	for _, d := range b.Defines {
		defs = append(defs, "-D"+d)
	}
	inc := []string{
		"-I" + j.SketchDir,
		"-I" + filepath.Join(coreDir, "avr"),
		"-I" + filepath.Join(coreDir, "stm32"),
		"-I" + filepath.Join(coreDir, "api/deprecated"),
		"-I" + srcWrapInc,
		"-I" + filepath.Join(srcWrapInc, "LL"),
		"-I" + pkg.subRoot(v["halInc"]),
		"-I" + pkg.subRoot(v["halSrc"]),
		"-I" + pkg.subRoot(v["sysFam"]),
		"-I" + pkg.subRoot(v["usbDevInc"]),
		"-I" + pkg.subRoot(v["usbDevSrc"]),
		"-I" + pkg.subRoot(v["virtioInc"]),
		"-I" + pkg.subRoot(v["openampInc"]),
		"-I" + pkg.subRoot(v["openampInc2"]),
		"-I" + pkg.subRoot(v["libmetalInc"]),
		"-I" + pkg.subRoot(v["virtDriverInc"]),
		"-I" + cmsis,
		"-I" + pkg.subRoot(v["familyInc"]),
		"-I" + pkg.subRoot(v["deviceSrc"]),
		"-I" + pkg.subRoot(v["deviceSrcGcc"]),
		"-I" + filepath.Join(cmsisDsp, "Include"),
		"-I" + filepath.Join(cmsisDsp, "PrivateInclude"),
		"-I" + coreDir,
		"-I" + variantDir,
		"-I" + srcWrap,
	}
	for _, li := range libIncs {
		inc = append(inc, "-I"+li)
	}
	base := []string{
		v["cpuFlag"], v["vectTab"],
	}
	for _, d := range strings.Fields(v["halDrivers"]) {
		base = append(base, d)
	}
	base = append(base, "-mthumb", "@"+buildOpt, "-c", "-Os", "-DNDEBUG", "-w",
		"-ffunction-sections", "-fdata-sections", "-MMD")
	cflags := append(append(append([]string{}, base...), "-std=gnu17",
		"--param", "max-inline-insns-single=500"), defs...)
	cxxflags := append(append(append([]string{}, base...), "-std=gnu++17",
		"-fno-threadsafe-statics", "--param", "max-inline-insns-single=500",
		"-fno-rtti", "-fno-exceptions", "-fno-use-cxa-atexit"), defs...)
	sflags := append(append(append([]string{}, v["cpuFlag"], v["vectTab"]), "-mthumb",
		"@"+buildOpt, "-c", "-x", "assembler-with-cpp", "-w", "-MMD"), defs...)
	cflags = append(cflags, inc...)
	cxxflags = append(cxxflags, inc...)
	sflags = append(sflags, inc...)

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
		case ".s", ".S":
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
		v["cpuFlag"], "-mthumb", "-Os", "-DNDEBUG", "--specs=nano.specs",
		"-Wl,--defsym=LD_FLASH_OFFSET=" + v["flashOffset"],
		fmt.Sprintf("-Wl,--defsym=LD_MAX_SIZE=%d", b.FlashMax),
		fmt.Sprintf("-Wl,--defsym=LD_MAX_DATA_SIZE=%d", b.RamMax),
		"-Wl,--cref", "-Wl,--check-sections", "-Wl,--gc-sections",
		"-Wl,--entry=Reset_Handler", "-Wl,--unresolved-symbols=report-all",
		"-Wl,--warn-common",
		"-Wl,--default-script=" + pkg.subRoot(v["ldscript"]),
		"-Wl,--script=" + pkg.subRoot(v["ldscriptSys"]),
		"-Wl,-Map," + filepath.Join(j.OutDir, "firmware.map"),
		"-Wl,--no-warn-rwx-segments",
		"-o", elf,
		"-L" + j.OutDir,
		"-Wl,--start-group",
	}
	link = append(link, sketchObjs...)
	link = append(link, coreA, "-lc", "-Wl,--end-group", "-lm", "-lgcc", "-lstdc++")
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cc, Args: link})

	bin := filepath.Join(j.OutDir, "firmware.bin")
	hex := filepath.Join(j.OutDir, "firmware.hex")
	p.Steps = append(p.Steps, Step{Name: "bin", Exe: objcopy, Args: []string{"-O", "binary", elf, bin}})
	p.Steps = append(p.Steps, Step{Name: "hex", Exe: objcopy, Args: []string{"-O", "ihex", elf, hex}})
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}
