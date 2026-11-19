package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// esp32Steps builds the family:esp32 recipe instance. The ESP32 core 3.x is
// ESP-IDF based: compile flags ship as response files under <sdk>/flags, the
// prebuilt IDF libraries are linked as one group, and firmware lands as a .bin
// via esptool elf2image plus a merged image with bootloader+partitions.
// Board-specific values arrive via manifest vars; the sdk is the per-chip
// esp32<target>-libs package.
func (p *Plan) esp32Steps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	v := p.Vars
	for _, k := range []string{
		"tarch", "target", "mcu", "boardDefine", "fCpu",
		"boot", "bootFreq", "flashMode", "flashFreq", "flashSize",
		"psramType", "partitions", "usbMode", "cdcOnBoot", "mscOnBoot", "dfuOnBoot",
		"bootloaderAddr",
	} {
		if v[k] == "" {
			return fmt.Errorf("%s: esp32 manifest needs var %q", b.Short, k)
		}
	}
	arch := v["arch"]
	if arch == "" {
		arch = "ESP32"
	}
	ideVer := v["ideVersion"]
	if ideVer == "" {
		ideVer = "10819"
	}
	dbg := v["codeDebug"]
	if dbg == "" {
		dbg = "0"
	}
	memType := v["boot"] + "_" + v["psramType"] // e.g. qio_qspi

	sdk := ""
	if pkg.SDKs != nil {
		sdk = pkg.SDKs[v["mcu"]]
	}
	if sdk == "" {
		return &unsupported{fmt.Sprintf("package %s has no SDK for %s", pkg.ID, v["mcu"])}
	}
	if st, err := os.Stat(filepath.Join(sdk, "flags")); err != nil || !st.IsDir() {
		return &unsupported{fmt.Sprintf("sdk for %s not installed at %s", v["mcu"], sdk)}
	}

	// Xtensa chips build with their own prefixed binaries (xtensa-esp32s3-elf-*)
	// under gccBin; the riscv chips' generic riscv32-esp-elf-* set (their
	// boards.txt target is "esp") lives in a separate package dir.
	gccBin := pkg.Tools["gccBin"]
	if v["tarch"] == "riscv32" {
		gccBin = pkg.Tools["riscvBin"]
	}
	if gccBin == "" {
		return &unsupported{fmt.Sprintf("package %s has no %s tool dir", pkg.ID, map[bool]string{true: "riscvBin", false: "gccBin"}[v["tarch"] == "riscv32"])}
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
	ar, ok3 := exeAt(prefix + "gcc-ar")
	size, ok4 := exeAt(prefix + "size")
	if !(ok1 && ok2 && ok3 && ok4) {
		return &unsupported{fmt.Sprintf("xtensa/riscv toolchain for %s missing under %s", v["target"], gccBin)}
	}
	esptool, ok := pkg.tool("esptool")
	if !ok {
		return &unsupported{fmt.Sprintf("esptool missing from package %s", pkg.ID)}
	}
	genpart, ok := pkg.tool("genpart")
	if !ok {
		return &unsupported{fmt.Sprintf("gen_esp32part missing from package %s", pkg.ID)}
	}

	coreDir := pkg.subRoot(b.Core)
	variantDir := pkg.subRoot(b.Variant)
	partsDir := pkg.subRoot("tools/partitions")

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

	// Files the recipes consume as @response files: an empty build_opt.h (sketch may
	// supply one), file_opts carrying -DARDUINO_CORE_BUILD for core units only, and
	// the partition table CSV (sketch override wins).
	if err := os.MkdirAll(j.OutDir, 0o755); err != nil {
		return err
	}
	partsCSV := filepath.Join(j.OutDir, "partitions.csv")
	if err := copyOr(filepath.Join(j.SketchDir, "partitions.csv"), partsCSV); err != nil {
		if err := copyOr(filepath.Join(partsDir, v["partitions"]+".csv"), partsCSV); err != nil {
			return fmt.Errorf("partitions %q: %w", v["partitions"], err)
		}
	}
	buildOpt := filepath.Join(j.OutDir, "build_opt.h")
	if err := copyOr(filepath.Join(j.SketchDir, "build_opt.h"), buildOpt); err != nil {
		if err := os.WriteFile(buildOpt, nil, 0o644); err != nil {
			return err
		}
	}
	fileOpts := filepath.Join(j.OutDir, "file_opts.txt")
	fileOptsCore := filepath.Join(j.OutDir, "file_opts_core.txt")
	if err := os.WriteFile(fileOpts, nil, 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(fileOptsCore, []byte("-DARDUINO_CORE_BUILD"), 0o644); err != nil {
		return err
	}

	sdkFlag := func(name string) string { return "@" + filepath.Join(sdk, "flags", name) }
	fqbn := b.ID
	var optKeys []string
	for k := range p.Options {
		optKeys = append(optKeys, k)
	}
	sort.Strings(optKeys)
	for _, k := range optKeys {
		fqbn += "." + k + "=" + p.Options[k]
	}
	defs := []string{
		"-DF_CPU=" + v["fCpu"],
		"-DARDUINO=" + ideVer,
		"-DARDUINO_" + v["boardDefine"],
		"-DARDUINO_ARCH_" + arch,
		`-DARDUINO_BOARD="` + v["boardDefine"] + `"`,
		`-DARDUINO_VARIANT="` + filepath.Base(b.Variant) + `"`,
		"-DARDUINO_PARTITION_" + v["partitions"],
		`-DARDUINO_HOST_OS="windows"`,
		`-DARDUINO_FQBN="` + fqbn + `"`,
		"-DESP32",
		"-DCORE_DEBUG_LEVEL=" + dbg,
		"-DARDUINO_USB_MODE=" + v["usbMode"],
		"-DARDUINO_USB_CDC_ON_BOOT=" + v["cdcOnBoot"],
		"-DARDUINO_USB_MSC_ON_BOOT=" + v["mscOnBoot"],
		"-DARDUINO_USB_DFU_ON_BOOT=" + v["dfuOnBoot"],
	}
	for _, d := range b.Defines {
		defs = append(defs, "-D"+d)
	}
	inc := []string{
		sdkFlag("defines"),
		"-I" + j.SketchDir,
		"-iprefix", filepath.Join(sdk, "include") + string(filepath.Separator),
		sdkFlag("includes"),
		"-I" + filepath.Join(sdk, memType, "include"),
		"-I" + coreDir,
		"-I" + variantDir,
		"-I" + filepath.Join(j.OutDir, "preproc"),
	}
	for _, li := range libIncs {
		inc = append(inc, "-I"+li)
	}
	cflags := append(append([]string{"-MMD", "-c", sdkFlag("c_flags"), "-w", "-Os", "-Werror=return-type"}, defs...), inc...)
	cxxflags := append(append([]string{"-MMD", "-c", sdkFlag("cpp_flags"), "-w", "-Os", "-Werror=return-type"}, defs...), inc...)
	sflags := append(append([]string{"-MMD", "-c", "-x", "assembler-with-cpp", sdkFlag("S_flags"), "-w", "-Os"}, defs...), inc...)

	// Core objects/archive land in the shared warm cache when enabled — an ESP32
	// core build is the slowest part of a cold build, so reuse matters most here.
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
	add := func(root, src, opts string) string {
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
		args = append(args, "@"+buildOpt, "@"+opts, src, "-o", obj+".part")
		p.Steps = append(p.Steps, Step{Name: "compile " + filepath.Base(src), Exe: compilerFor(src, cc, cxx), Parallel: true, Out: obj, Args: args})
		return obj
	}

	var coreObjs []string
	if !warm {
		for _, src := range coreSrcs {
			coreObjs = append(coreObjs, add(coreObjRoot, src, fileOptsCore))
		}
		p.Steps = append(p.Steps, Step{Name: "archive core", Exe: ar, Args: append([]string{"cr", coreA}, coreObjs...)})
	}
	var sketchObjs []string
	for _, src := range sketchSrcs {
		sketchObjs = append(sketchObjs, add(filepath.Join(objRoot, "sketch"), src, fileOpts))
	}
	for name, srcs := range libSrcs {
		for _, src := range srcs {
			sketchObjs = append(sketchObjs, add(filepath.Join(objRoot, "libs", name), src, fileOpts))
		}
	}

	elf := filepath.Join(j.OutDir, "firmware.elf")
	link := []string{
		"-Wl,--Map=" + filepath.Join(j.OutDir, "firmware.map"),
		"-L" + filepath.Join(sdk, "lib"),
		"-L" + filepath.Join(sdk, "ld"),
		"-L" + filepath.Join(sdk, memType),
		"-Wl,--wrap=esp_panic_handler",
		"-Wl,--wrap=esp_bt_mem_release",
		"-Wl,--wrap=esp_bt_controller_mem_release",
		"-Wl,--wrap=_ZN4chip11DeviceLayer8Internal10ESP32Utils13InitWiFiStackEv",
		"-Wl,--wrap=_Z21esp_matter_mem_callocjj",
		"-Wl,--wrap=_Z22esp_matter_mem_reallocPvj",
		"-Wl,--wrap=_Z19esp_matter_mem_freePv",
		sdkFlag("ld_flags"),
		sdkFlag("ld_scripts"),
		"-Wl,--start-group",
	}
	link = append(link, sketchObjs...)
	link = append(link, coreA, sdkFlag("ld_libs"), "-Wl,--end-group", "-Wl,-EL", "-o", elf)
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cxx, Args: link})

	bootBin := filepath.Join(j.OutDir, "bootloader.bin")
	partsBin := filepath.Join(j.OutDir, "firmware.partitions.bin")
	artifact := filepath.Join(j.OutDir, "firmware.bin")
	merged := filepath.Join(j.OutDir, "firmware.merged.bin")
	img := []string{"--chip", v["mcu"], "elf2image", "--flash-mode", v["flashMode"], "--flash-freq", v["flashFreq"], "--flash-size", v["flashSize"]}
	p.Steps = append(p.Steps, Step{Name: "bootloader", Exe: esptool,
		Args: append(append([]string{}, img...), "-o", bootBin, filepath.Join(sdk, "bin", "bootloader_"+v["boot"]+"_"+v["bootFreq"]+".elf"))})
	p.Steps = append(p.Steps, Step{Name: "partitions", Exe: genpart, Args: []string{"-q", partsCSV, partsBin}})
	p.Steps = append(p.Steps, Step{Name: "bin", Exe: esptool,
		Args: append(append([]string{}, img...), "--elf-sha256-offset", "0xb0", "-o", artifact, elf)})
	p.Steps = append(p.Steps, Step{Name: "merged", Exe: esptool,
		Args: []string{"--chip", v["mcu"], "merge-bin", "-o", merged, "--pad-to-size", v["flashSize"],
			"--flash-mode", "keep", "--flash-freq", "keep", "--flash-size", "keep",
			v["bootloaderAddr"], bootBin, "0x8000", partsBin, "0xe000",
			filepath.Join(partsDir, "boot_app0.bin"), "0x10000", artifact}})
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}

// sourcesRecursive lists .c/.cpp/.s sources under dir at any depth — ESP32 cores
// keep code in subdirs (libb64) that top-level listing would miss.
func sourcesRecursive(dir string) ([]string, error) {
	var out []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if srcExts[strings.ToLower(filepath.Ext(info.Name()))] {
			out = append(out, path)
		}
		return nil
	})
	return out, err
}

// copyOr copies src to dst when src exists; os.ErrNotExist means "leave dst alone"
// so callers can fall back to a generated default.
func copyOr(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}

// compilerFor picks cc or cxx for one source path by extension.
func compilerFor(src, cc, cxx string) string {
	switch strings.ToLower(filepath.Ext(src)) {
	case ".cpp", ".cxx", ".cc":
		return cxx
	}
	return cc
}
