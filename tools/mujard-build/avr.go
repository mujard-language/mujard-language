package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// avrSteps builds the family:avr recipe instance: compile core, compile sketch,
// archive, link, objcopy, size. Flag shapes come from the AVR recipe the old platform
// encoded in platform.txt; every value that differs per board arrives via manifest vars.
func (p *Plan) avrSteps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	mcu := p.Vars["mcu"]
	fcpu := p.Vars["fCpu"]
	bdef := p.Vars["boardDefine"]
	aver := p.Vars["arduinoVersion"]
	if mcu == "" || fcpu == "" || bdef == "" {
		return fmt.Errorf("%s: avr manifest needs vars mcu, fCpu and boardDefine", b.Short)
	}
	if aver == "" {
		aver = "10813"
	}
	coreDir := pkg.subRoot(b.Core)
	variantDir := pkg.subRoot(b.Variant)

	coreSrcs, _, err := sources(coreDir)
	if err != nil {
		return fmt.Errorf("core sources: %w", err)
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

	for _, t := range []string{"cc", "cxx", "ar", "objcopy", "size"} {
		if _, ok := pkg.tool(t); !ok {
			return &unsupported{fmt.Sprintf("tool %q missing from package %s", t, pkg.ID)}
		}
	}
	cc, _ := pkg.tool("cc")
	cxx, _ := pkg.tool("cxx")
	ar, _ := pkg.tool("ar")
	objcopy, _ := pkg.tool("objcopy")
	size, _ := pkg.tool("size")

	defines := []string{
		"-DF_CPU=" + fcpu,
		"-DARDUINO=" + aver,
		"-D" + bdef,
		"-DARDUINO_ARCH_AVR",
	}
	for _, d := range b.Defines {
		defines = append(defines, "-D"+d)
	}
	inc := []string{"-I" + coreDir, "-I" + variantDir, "-I" + j.SketchDir, "-I" + filepath.Join(j.OutDir, "preproc")}
	for _, li := range libIncs {
		inc = append(inc, "-I"+li)
	}
	common := []string{"-c", "-g", "-Os", "-w", "-ffunction-sections", "-fdata-sections", "-MMD", "-flto", "-mmcu=" + mcu}
	cflags := append(append(append([]string{"-std=gnu11", "-fno-fat-lto-objects"}, common...), defines...), inc...)
	cxxflags := append(append(append([]string{"-std=gnu++11", "-fpermissive", "-fno-exceptions", "-fno-threadsafe-statics", "-Wno-error=narrowing"}, common...), defines...), inc...)
	sflags := append(append(append([]string{"-x", "assembler-with-cpp"}, common...), defines...), inc...)

	// Where the core's objects and archive land: the shared cache when it is cold
	// (staging, published on success), the published cache when warm, or the job folder
	// when caching is off.
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

	var coreObjs, sketchObjs []string
	objs := objNamer{}
	addCompile := func(root, src string) string {
		obj := objs.path(root, src)
		p.Steps = append(p.Steps, compileStep(cc, cxx, cflags, cxxflags, sflags, src, obj))
		return obj
	}
	if !warm {
		for _, src := range coreSrcs {
			coreObjs = append(coreObjs, addCompile(coreObjRoot, src))
		}
		p.Steps = append(p.Steps, Step{Name: "archive core", Exe: ar, Args: append([]string{"rcs", coreA}, coreObjs...)})
	}
	for _, src := range sketchSrcs {
		sketchObjs = append(sketchObjs, addCompile(filepath.Join(objRoot, "sketch"), src))
	}
	for name, srcs := range libSrcs {
		for _, src := range srcs {
			sketchObjs = append(sketchObjs, addCompile(filepath.Join(objRoot, "libs", name), src))
		}
	}

	elf := filepath.Join(j.OutDir, "firmware.elf")
	link := append([]string{"-w", "-Os", "-g", "-flto", "-fuse-linker-plugin", "-Wl,--gc-sections", "-mmcu=" + mcu, "-o", elf}, sketchObjs...)
	link = append(link, coreA, "-L"+j.OutDir, "-lm")
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cc, Args: link})

	var artifact string
	switch b.Output {
	case "hex":
		artifact = filepath.Join(j.OutDir, "firmware.hex")
		p.Steps = append(p.Steps, Step{Name: "hex", Exe: objcopy, Args: []string{"-O", "ihex", "-R", ".eeprom", elf, artifact}})
	default:
		return &unsupported{fmt.Sprintf("output %q has no image step yet", b.Output)}
	}
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}

// objNamer gives every source its own object file. The name keeps the source's extension, as arduino-cli
// does: the AVR core has both wiring_pulse.c and wiring_pulse.S, and when both became wiring_pulse.o one
// overwrote the other, so pulseIn() failed to link (countPulseASM). A second source with the same file
// name under one folder (a library's x.c and utility/x.c) gets a number. Names compare case-insensitively
// because Windows file names do.
type objNamer map[string]bool

func (n objNamer) path(root, src string) string {
	base := filepath.Base(src)
	obj := filepath.Join(root, base+".o")
	for i := 2; n[strings.ToLower(obj)]; i++ {
		obj = filepath.Join(root, fmt.Sprintf("%s.%d.o", base, i))
	}
	n[strings.ToLower(obj)] = true
	return obj
}

// compileStep picks the right compiler+flags for one source's extension.
func compileStep(cc, cxx string, cflags, cxxflags, sflags []string, src, obj string) Step {
	name := "compile " + filepath.Base(src)
	switch strings.ToLower(filepath.Ext(src)) {
	case ".c":
		return Step{Name: name, Exe: cc, Parallel: true, Out: obj, Args: append(append([]string{}, cflags...), "-o", obj+".part", src)}
	case ".s":
		return Step{Name: name, Exe: cc, Parallel: true, Out: obj, Args: append(append([]string{}, sflags...), "-o", obj+".part", src)}
	default:
		return Step{Name: name, Exe: cxx, Parallel: true, Out: obj, Args: append(append([]string{}, cxxflags...), "-o", obj+".part", src)}
	}
}
