package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// rp2Steps builds the family:rp2 recipe instance: the Raspberry Pi RP2040 and RP2350 with
// the arduino-pico core. It follows that core's platform.txt: part of the compile flags ship
// as response files under lib/, the prebuilt Pico SDK (libpico.a, lwIP, BearSSL) is linked
// as one group, the linker script is the chip's memmap template with this board's sizes
// filled in, and picotool turns the ELF into a UF2. The core runs Python for the two small
// steps before the link (simplesub.py, signing.py --mode header); both are done here so a
// pack needs no Python.
func (p *Plan) rp2Steps() error {
	b, j, pkg := p.Board, p.Job, p.Pkg
	v := p.Vars
	for _, k := range []string{
		"chip", "toolchainOpts", "boardDefine", "fCpu", "usbVid", "usbPid", "usbManufacturer", "usbProduct",
		"flashTotal", "flashLength", "eepromStart", "fsStart", "fsEnd", "ramLength", "psramLength", "boot2", "uf2Family",
	} {
		if v[k] == "" {
			return fmt.Errorf("%s: rp2 manifest needs var %q", b.Short, k)
		}
	}
	ideVer := v["ideVersion"]
	if ideVer == "" {
		ideVer = "10607"
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
	if !(ok1 && ok2 && ok3 && ok4) {
		return &unsupported{fmt.Sprintf("arm-none-eabi toolchain missing under %s", gccBin)}
	}
	picotool, ok := pkg.tool("picotool")
	if !ok {
		return &unsupported{fmt.Sprintf("picotool missing from package %s", pkg.ID)}
	}

	root := pkg.Root
	chip := v["chip"]
	libDir := filepath.Join(root, "lib", chip)
	memmapIn := filepath.Join(libDir, "memmap_default.ld")
	if _, err := os.Stat(memmapIn); err != nil {
		return &unsupported{fmt.Sprintf("the %s core files are not installed at %s", chip, root)}
	}
	coreDir := pkg.subRoot(b.Core)
	variantDir := pkg.subRoot(b.Variant)

	coreSrcs, err := sourcesRecursive(coreDir)
	if err != nil {
		return fmt.Errorf("core sources: %w", err)
	}
	if varSrcs, err := sourcesRecursive(variantDir); err == nil {
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

	// signing.py --mode header without a public.key: the core's Updater is built unsigned.
	genCore := filepath.Join(j.OutDir, "core")
	if err := os.MkdirAll(genCore, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(genCore, "Updater_Signing.h"), []byte("#define ARDUINO_SIGNING 0\n"), 0o644); err != nil {
		return err
	}
	// simplesub.py: the chip's linker script template with this board's flash/RAM layout.
	memmap := filepath.Join(j.OutDir, "memmap_default.ld")
	tmpl, err := os.ReadFile(memmapIn)
	if err != nil {
		return err
	}
	ld := strings.NewReplacer(
		"__FLASH_LENGTH__", v["flashLength"], "__EEPROM_START__", v["eepromStart"],
		"__FS_START__", v["fsStart"], "__FS_END__", v["fsEnd"],
		"__RAM_LENGTH__", v["ramLength"], "__PSRAM_LENGTH__", v["psramLength"],
	).Replace(string(tmpl))
	if err := os.WriteFile(memmap, []byte(ld), 0o644); err != nil {
		return err
	}

	opts := strings.Fields(v["toolchainOpts"])
	warn := []string{"-Werror=return-type", "-Wno-psabi"}
	defs := []string{
		"-DUSBD_PID=" + v["usbPid"], "-DUSBD_VID=" + v["usbVid"], "-DUSBD_MAX_POWER_MA=250",
		`-DUSB_MANUFACTURER="` + v["usbManufacturer"] + `"`, `-DUSB_PRODUCT="` + v["usbProduct"] + `"`,
		"-DLWIP_IPV6=0", "-DLWIP_IPV4=1", "-DLWIP_IGMP=1", "-DLWIP_CHECKSUM_CTRL_PER_NETIF=1",
		"-DFILE_COPY_CONSTRUCTOR_SELECT=FILE_COPY_CONSTRUCTOR_PUBLIC", "-DUSE_UTF8_LONG_NAMES=1", "-DDISABLE_FS_H_WARNING=1",
	}
	for _, d := range b.Defines {
		defs = append(defs, "-D"+d)
	}
	defs = append(defs,
		`-DARDUINO_VARIANT="`+filepath.Base(b.Variant)+`"`,
		"-DPICO_FLASH_SIZE_BYTES="+v["flashTotal"], "-DFS_START="+v["fsStart"], "-DFS_END="+v["fsEnd"],
		"@"+filepath.Join(root, "lib", "platform_def.txt"), "@"+filepath.Join(libDir, "platform_def.txt"),
	)
	common := append(append([]string{}, opts...), "-ffunction-sections", "-fdata-sections", "-fno-exceptions")
	inc := []string{
		"-iprefix" + filepath.ToSlash(root) + "/",
		"@" + filepath.Join(libDir, "platform_inc.txt"),
		"@" + filepath.Join(root, "lib", "core_inc.txt"),
		"-I" + filepath.Join(root, "include"),
	}
	board := []string{
		"-DF_CPU=" + v["fCpu"], "-DARDUINO=" + ideVer, "-DARDUINO_" + v["boardDefine"],
		`-DBOARD_NAME="` + v["boardDefine"] + `"`, "-DARDUINO_ARCH_RP2040",
	}
	join := func(parts ...[]string) []string {
		var out []string
		for _, p := range parts {
			out = append(out, p...)
		}
		return out
	}
	// The library scan follows #ifdef blocks with the macros these flags (and the core's platform_def.txt files) define.
	p.Defined = map[string]bool{"WIFICC": true}
	for _, f := range join(defs, board) {
		if strings.HasPrefix(f, "@") {
			if data, err := os.ReadFile(f[1:]); err == nil {
				for _, w := range strings.Fields(string(data)) {
					if strings.HasPrefix(w, "-D") {
						p.Defined[strings.SplitN(w[2:], "=", 2)[0]] = true
					}
				}
			}
		} else if strings.HasPrefix(f, "-D") {
			p.Defined[strings.SplitN(f[2:], "=", 2)[0]] = true
		}
	}
	libIncs, libSrcs, err := p.resolveLibs()
	if err != nil {
		return fmt.Errorf("libraries: %w", err)
	}

	// The core is compiled with its own folders only, as arduino-cli does; libraries are for the sketch and themselves.
	coreInc := []string{"-I" + coreDir, "-I" + variantDir}
	srcInc := append(append([]string{}, coreInc...), "-I"+j.SketchDir, "-I"+filepath.Join(j.OutDir, "preproc"))
	for _, li := range libIncs {
		srcInc = append(srcInc, "-I"+li)
	}
	flagsFor := func(incs []string) (c, cxx, s []string) {
		c = join([]string{"-c"}, warn, defs, common, []string{"-MMD"}, inc, []string{"-std=gnu23", "-g", "-pipe"}, board, []string{"-Os"}, incs)
		cxx = join([]string{"-I" + genCore, "-c"}, warn, defs, common, []string{"-MMD"}, inc,
			[]string{"-fno-rtti", "-std=gnu++23", "-g", "-pipe", "-Wno-volatile"}, board, []string{"-Os", "-DWIFICC=CYW43_COUNTRY_WORLDWIDE"}, incs)
		s = join([]string{"-c"}, warn, defs, []string{"-g", "-x", "assembler-with-cpp", "-MMD"}, inc, opts, []string{"-g"}, board, incs)
		return
	}
	coreC, coreCxx, coreS := flagsFor(coreInc)
	cflags, cxxflags, sflags := flagsFor(srcInc)
	elfFlags := join(warn, defs, common, []string{"-Os", "-u", "_printf_float", "-u", "_scanf_float"})

	objRoot := filepath.Join(j.OutDir, "obj")
	coreA := filepath.Join(j.OutDir, "core.a")
	coreObjRoot := filepath.Join(objRoot, "core")
	warm := false
	if !j.NoCache {
		if final := p.coreCacheDir(coreC, coreCxx, coreS); final != "" {
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
	addCompile := func(root, src string) string {
		obj := objs.path(root, src)
		p.Steps = append(p.Steps, compileStep(cc, cxx, cflags, cxxflags, sflags, src, obj))
		return obj
	}
	if !warm {
		var coreObjs []string
		for _, src := range coreSrcs {
			obj := objs.path(coreObjRoot, src)
			p.Steps = append(p.Steps, compileStep(cc, cxx, coreC, coreCxx, coreS, src, obj))
			coreObjs = append(coreObjs, obj)
		}
		p.Steps = append(p.Steps, Step{Name: "archive core", Exe: ar, Args: append([]string{"rcs", coreA}, coreObjs...)})
	}
	var sketchObjs []string
	for _, src := range sketchSrcs {
		sketchObjs = append(sketchObjs, addCompile(filepath.Join(objRoot, "sketch"), src))
	}
	for name, srcs := range libSrcs {
		for _, src := range srcs {
			sketchObjs = append(sketchObjs, addCompile(filepath.Join(objRoot, "libs", name), src))
		}
	}

	boot2 := filepath.Join(j.OutDir, "boot2.o")
	p.Steps = append(p.Steps, Step{Name: "boot2", Exe: cc, Args: join(elfFlags, []string{
		"-c", filepath.Join(root, "boot2", chip, v["boot2"]+".S"),
		"-I" + filepath.Join(root, "pico-sdk", "src", chip, "hardware_regs", "include"),
		"-I" + filepath.Join(root, "pico-sdk", "src", "common", "pico_binary_info", "include"),
		"-o", boot2,
	})})

	elf := filepath.Join(j.OutDir, "firmware.elf")
	link := join([]string{"-L" + j.OutDir}, elfFlags, []string{
		"@" + filepath.Join(libDir, "platform_wrap.txt"), "@" + filepath.Join(root, "lib", "core_wrap.txt"),
		"-Wl,--cref", "-Wl,--check-sections", "-Wl,--gc-sections", "-Wl,--unresolved-symbols=report-all", "-Wl,--warn-common",
	})
	for _, u := range []string{
		"runtime_init_install_ram_vector_table", "__pre_init_runtime_init_clocks", "__pre_init_runtime_init_bootrom_reset",
		"__pre_init_runtime_init_early_resets", "__pre_init_runtime_init_usb_power_down", "__pre_init_runtime_init_post_clock_resets",
		"__pre_init_runtime_init_spin_locks_reset", "__pre_init_runtime_init_boot_locks_reset",
		"__pre_init_runtime_init_bootrom_locking_enable", "__pre_init_runtime_init_mutex",
		"__pre_init_runtime_init_default_alarm_pool", "__pre_init_first_per_core_initializer",
		"__pre_init_runtime_init_per_core_bootrom_reset", "__pre_init_runtime_init_per_core_h3_irq_registers",
		"__pre_init_runtime_init_per_core_irq_priorities",
	} {
		link = append(link, "-Wl,--undefined="+u)
	}
	link = append(link, "-Wl,--script="+memmap, "-Wl,-Map,"+filepath.Join(j.OutDir, "firmware.map"),
		"-o", elf, "-Wl,--no-warn-rwx-segments", "-Wl,--start-group")
	link = append(link, sketchObjs...)
	link = append(link, coreA, boot2, filepath.Join(libDir, "ota.o"),
		filepath.Join(libDir, "libpico.a"), filepath.Join(libDir, "liblwip.a"), filepath.Join(libDir, "libbearssl.a"),
		"-lm", "-lc", "-lstdc++", "-lc", "-Wl,--end-group")
	p.Steps = append(p.Steps, Step{Name: "link", Exe: cxx, Args: link})

	switch b.Output {
	case "uf2":
		uf2 := filepath.Join(j.OutDir, "firmware.uf2")
		p.Steps = append(p.Steps, Step{Name: "uf2", Exe: picotool,
			Args: append([]string{"uf2", "convert", elf, uf2}, strings.Fields(v["uf2Family"])...)})
	default:
		return &unsupported{fmt.Sprintf("output %q has no image step for rp2", b.Output)}
	}
	p.Steps = append(p.Steps, Step{Name: "size", Exe: size, Args: []string{"-A", elf}})
	return nil
}
