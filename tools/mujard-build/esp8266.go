package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// esp8266Steps builds the family:esp8266 recipe instance. The ESP8266 core is Arduino-style
// (like AVR) but xtensa-lx106: compile core + sketch, archive, link against the SDK's
// eagle ld scripts (preprocessed with the vtable/MMU defines), then the core's own
// elf2bin.py produces the flash image around the eboot bootloader. Flag shapes come from
// the core's platform.txt recipe, proven against an arduino-cli verbose trace.
func (p *Plan) esp8266Steps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	v := p.Vars
	for _, k := range []string{
		"target", "boardDefine", "boardId", "fCpu", "sdk",
		"lwipInclude", "lwipLib", "lwipFlags", "led",
		"vtable", "mmuIram", "mmuIcache", "exceptionFlags", "stdcppLib",
		"floatFlags", "flashLd", "flashMode", "flashFreq", "flashSize", "flashFlag",
	} {
		if v[k] == "" {
			return fmt.Errorf("%s: esp8266 manifest needs var %q", b.Short, k)
		}
	}
	arch := v["arch"]
	if arch == "" {
		arch = "ESP8266"
	}
	ideVer := v["ideVersion"]
	if ideVer == "" {
		ideVer = "10607"
	}

	gccBin := pkg.Tools["gccBin"]
	if gccBin == "" {
		return &unsupported{fmt.Sprintf("package %s has no gccBin tool dir", pkg.ID)}
	}
	prefix := v["tarch"] + "-" + v["target"] + "-elf-"
	exeAt := func(name string) (string, bool) {
		for _, c := range []string{filepath.Join(gccBin, name), filepath.Join(gccBin, name+".exe")} {
			if st, err := os.Stat(c); err == nil && !st.IsDir() {
				return c, true
			}
		}
		return "", false
	}
	cc, ok1 := exeAt(prefix + "gcc")
	cxx, ok2 := exeAt(prefix + "g++")
	ar, ok3 := exeAt(prefix + "ar")
	size, ok4 := exeAt(prefix + "size")
	if !(ok1 && ok2 && ok3 && ok4) {
		return &unsupported{fmt.Sprintf("xtensa-lx106 toolchain missing under %s", gccBin)}
	}
	python, ok := pkg.tool("python3")
	if !ok {
		return &unsupported{fmt.Sprintf("python3 missing from package %s", pkg.ID)}
	}
	gccInc := filepath.Join(filepath.Dir(gccBin), "include")

	coreDir := pkg.subRoot(b.Core)
	variantDir := pkg.subRoot(b.Variant)
	sdkDir := pkg.subRoot("tools/sdk")
	libcDir := pkg.subRoot("tools/sdk/libc/xtensa-lx106-elf")
	warnC := pkg.subRoot("tools/warnings/none-cflags")
	warnCpp := pkg.subRoot("tools/warnings/none-cppflags")
	elf2bin := pkg.subRoot("tools/elf2bin.py")
	eboot := pkg.subRoot("bootloaders/eboot/eboot.elf")
	for _, need := range []string{coreDir, sdkDir, warnC, warnCpp, elf2bin, eboot} {
		if _, err := os.Stat(need); err != nil {
			return &unsupported{fmt.Sprintf("package %s missing %s", pkg.ID, need)}
		}
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

	// The recipe's @{build.path}/core/build.opt: sketch build_opt.h when present, else empty.
	coreOut := filepath.Join(j.OutDir, "core")
	if err := os.MkdirAll(coreOut, 0o755); err != nil {
		return err
	}
	buildOpt := filepath.Join(coreOut, "build.opt")
	if err := copyOr(filepath.Join(j.SketchDir, "build_opt.h"), buildOpt); err != nil {
		if err := os.WriteFile(buildOpt, nil, 0o644); err != nil {
			return err
		}
	}

	defs := []string{
		"-D__ets__", "-DICACHE_FLASH", "-U__STRICT_ANSI__", "-D_GNU_SOURCE", "-DESP8266",
		"@" + buildOpt,
		"-I" + filepath.Join(sdkDir, "include"),
		"-I" + filepath.Join(sdkDir, filepath.FromSlash(v["lwipInclude"])),
		"-I" + filepath.Join(libcDir, "include"),
		"-I" + coreOut,
	}
	defs = append(defs, strings.Fields(v["lwipFlags"])...)
	defs = append(defs,
		"-D"+v["sdk"]+"=1",
		"-DF_CPU="+v["fCpu"],
		"-DARDUINO="+ideVer,
		"-DARDUINO_"+v["boardDefine"],
		"-DARDUINO_ARCH_"+arch,
		`-DARDUINO_BOARD="`+v["boardDefine"]+`"`,
		`-DARDUINO_BOARD_ID="`+v["boardId"]+`"`,
		"-DLED_BUILTIN="+v["led"],
		v["flashFlag"],
	)
	for _, d := range b.Defines {
		defs = append(defs, "-D"+d)
	}
	inc := []string{"-I" + coreDir, "-I" + variantDir, "-I" + j.SketchDir, "-I" + filepath.Join(j.OutDir, "preproc")}
	for _, li := range libIncs {
		inc = append(inc, "-I"+li)
	}
	opt := append(defs, inc...)

	mmu := []string{"-DMMU_IRAM_SIZE=" + v["mmuIram"], "-DMMU_ICACHE_SIZE=" + v["mmuIcache"]}
	cflags := append(append([]string{
		"-c", "@" + warnC, "-std=gnu17",
		"-Os", "-g", "-free", "-fipa-pta", "-Werror=return-type", "-Wpointer-arith",
		"-Wno-implicit-function-declaration", "-Wl,-EL", "-fno-inline-functions",
		"-nostdlib", "-mlongcalls", "-mtext-section-literals", "-falign-functions=4",
		"-MMD", "-ffunction-sections", "-fdata-sections",
	}, strings.Fields(v["exceptionFlags"])...), mmu...)
	cflags = append(cflags, opt...)
	cxxflags := append(append([]string{
		"-c", "@" + warnCpp,
		"-Os", "-g", "-free", "-fipa-pta", "-Werror=return-type",
		"-mlongcalls", "-mtext-section-literals", "-fno-rtti", "-falign-functions=4",
		"-std=gnu++17", "-MMD", "-ffunction-sections", "-fdata-sections",
	}, strings.Fields(v["exceptionFlags"])...), mmu...)
	cxxflags = append(cxxflags, opt...)
	sflags := append(append([]string{
		"-c", "-g", "-x", "assembler-with-cpp", "-MMD", "-mlongcalls", "-I" + gccInc,
	}, mmu...), opt...)

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
	add := func(root, src string) string {
		obj := objs.path(root, src)
		p.Steps = append(p.Steps, compileStep(cc, cxx, cflags, cxxflags, sflags, src, obj))
		return obj
	}
	if !warm {
		for _, src := range coreSrcs {
			coreObjs = append(coreObjs, add(coreObjRoot, src))
		}
		p.Steps = append(p.Steps, Step{Name: "archive core", Exe: ar, Args: append([]string{"cru", coreA}, coreObjs...)})
	}
	for _, src := range sketchSrcs {
		sketchObjs = append(sketchObjs, add(filepath.Join(objRoot, "sketch"), src))
	}
	for name, srcs := range libSrcs {
		for _, src := range srcs {
			sketchObjs = append(sketchObjs, add(filepath.Join(objRoot, "libs", name), src))
		}
	}

	// Linker scripts are C-preprocessed from the SDK's .h templates with the vtable/MMU
	// defines (the core's prelink hooks do the same through cp.py + gcc -E).
	ldH := filepath.Join(j.OutDir, "ld_h")
	if err := os.MkdirAll(ldH, 0o755); err != nil {
		return err
	}
	flashLdH := filepath.Join(ldH, "local.eagle.flash.ld.h")
	if err := copyOr(filepath.Join(sdkDir, "ld", v["flashLd"]), flashLdH); err != nil {
		return fmt.Errorf("ld template %q: %w", v["flashLd"], err)
	}
	ldCflags := append(append([]string{"-CC", "-E", "-P"}, strings.Fields(v["vtable"])...), mmu...)
	flashLd := filepath.Join(j.OutDir, "local.eagle.flash.ld")
	commonLd := filepath.Join(j.OutDir, "local.eagle.app.v6.common.ld")
	p.Steps = append(p.Steps, Step{Name: "ld flash", Exe: cc, Args: append(append([]string{}, ldCflags...), flashLdH, "-o", flashLd)})
	p.Steps = append(p.Steps, Step{Name: "ld common", Exe: cc, Args: append(append([]string{}, ldCflags...), filepath.Join(sdkDir, "ld", "eagle.app.v6.common.ld.h"), "-o", commonLd)})

	elf := filepath.Join(j.OutDir, "firmware.elf")
	link := append(append([]string{}, strings.Fields(v["exceptionFlags"])...),
		"-Wl,-Map", "-Wl,"+filepath.Join(j.OutDir, "firmware.map"),
		"-g", "@"+warnC, "-Os", "-nostdlib",
		"-Wl,--no-check-sections", "-u", "app_entry",
	)
	link = append(link, strings.Fields(v["floatFlags"])...)
	link = append(link,
		"-Wl,-static",
		"-L"+filepath.Join(sdkDir, "lib"),
		"-L"+filepath.Join(sdkDir, "lib", v["sdk"]),
		"-L"+j.OutDir,
		"-L"+filepath.Join(libcDir, "lib"),
		"-T", flashLd,
		"-Wl,--gc-sections",
		"-Wl,-wrap,system_restart_local", "-Wl,-wrap,spi_flash_read",
		"-o", elf, "-Wl,--start-group",
	)
	link = append(link, sketchObjs...)
	link = append(link, coreA,
		"-lhal", "-lphy", "-lpp", "-lnet80211", v["lwipLib"], "-lwpa", "-lcrypto", "-lmain",
		"-lwps", "-lbearssl", "-lespnow", "-lsmartconfig", "-lairkiss", "-lwpa2",
		v["stdcppLib"], "-lm", "-lc", "-lgcc",
		"-Wl,--end-group", "-L"+j.OutDir,
	)
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cc, Args: link})

	artifact := filepath.Join(j.OutDir, "firmware.bin")
	p.Steps = append(p.Steps, Step{Name: "bin", Exe: python, Args: []string{
		"-I", elf2bin,
		"--eboot", eboot,
		"--app", elf,
		"--flash_mode", v["flashMode"],
		"--flash_freq", v["flashFreq"],
		"--flash_size", v["flashSize"],
		"--path", gccBin,
		"--out", artifact,
	}})
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}
