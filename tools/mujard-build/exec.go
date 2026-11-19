package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// runBuild executes a plan. A nil plan or *unsupported error means the caller should
// fall back to another backend — that is a result, not a crash.
func runBuild(job *Job, plan *Plan, perr error) *Result {
	res := &Result{Schema: SchemaV1}
	if perr != nil {
		var u *unsupported
		if errors.As(perr, &u) {
			res.Unsupported = true
		}
		res.Log = perr.Error()
		return res
	}
	res.Board = plan.Board.ID
	res.Format = plan.Board.Output

	start := time.Now()
	defer func() { res.Seconds = time.Since(start).Seconds() }()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(job.TimeoutSec)*time.Second)
	defer cancel()

	if err := os.MkdirAll(job.OutDir, 0o755); err != nil {
		res.Log = "cannot create outDir: " + err.Error()
		return res
	}

	var log strings.Builder
	var sizeOut string
	fail := func(out []byte, err error) *Result {
		if err != nil {
			res.Errors = parseDiags(string(out))
		}
		res.Log = log.String()
		if ctx.Err() == context.DeadlineExceeded {
			res.Log += "\nbuild timed out"
		}
		// A failed build publishes no cache, but staging survives: finished
		// objects let the next identical build resume where this one stopped.
		return res
	}
	for i := 0; i < len(plan.Steps); {
		s := plan.Steps[i]
		if s.Parallel {
			j := i
			for j < len(plan.Steps) && plan.Steps[j].Parallel {
				j++
			}
			outs, err := runParallel(ctx, job.OutDir, plan.Steps[i:j], &log)
			if ctx.Err() == context.DeadlineExceeded || err != nil {
				return fail([]byte(outs), err)
			}
			i = j
			continue
		}
		if err := os.MkdirAll(filepath.Dir(firstOutput(s.Args)), 0o755); err != nil {
			res.Log = log.String() + "mkdir: " + err.Error()
			return res
		}
		fmt.Fprintf(&log, "=== %s\n%s %s\n", s.Name, s.Exe, strings.Join(s.Args, " "))
		var out []byte
		var err error
		for attempt := 1; ; attempt++ {
			cmd := exec.CommandContext(ctx, s.Exe, s.Args...)
			cmd.Dir = job.OutDir
			out, err = cmd.CombinedOutput()
			log.Write(out)
			if err == nil || attempt >= maxToolStartAttempts || !toolFailedToStart(out) || ctx.Err() != nil {
				break
			}
			fmt.Fprintf(&log, "(%s could not unpack itself - retrying, attempt %d of %d)\n", filepath.Base(s.Exe), attempt+1, maxToolStartAttempts)
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
		}
		if s.Name == "size" {
			sizeOut = string(out)
		}
		if ctx.Err() == context.DeadlineExceeded || err != nil {
			return fail(out, err)
		}
		i++
	}
	res.Log = log.String()
	publishCore(plan)

	elf := filepath.Join(job.OutDir, "firmware.elf")
	artifact := filepath.Join(job.OutDir, "firmware."+plan.Board.Output)
	// ESP32-family: a merged image (bootloader + partitions + app) is the complete flashable image the
	// flasher writes at address 0; the bare firmware.bin is only the app partition and won't boot there.
	m := filepath.Join(job.OutDir, "firmware.merged.bin")
	if st, err := os.Stat(m); err == nil && st.Size() > 0 {
		artifact = m
	}
	if st, err := os.Stat(artifact); err != nil || st.Size() == 0 {
		res.Log = log.String() + "\nno " + plan.Board.Output + " produced"
		return res
	}
	res.Artifact = artifact
	res.Elf = elf
	res.Sizes = &Sizes{FlashMax: plan.Board.FlashMax, RamMax: plan.Board.RamMax}
	res.Sizes.Flash, res.Sizes.Ram = parseSize(sizeOut, plan.Board.Family)
	res.Success = true
	return res
}

// esptool.exe is a PyInstaller one-file program: every run first unpacks itself into a temp folder. On Windows an
// antivirus scan can lock a DLL it just wrote ("Failed to extract ucrtbase.dll ... Permission denied"), so the run
// fails before esptool did anything. That is retried; any other failure is the build's own and is not.
const maxToolStartAttempts = 3

var pyiExtractFailed = regexp.MustCompile(`\[PYI-\d+:ERROR\] Failed to extract`)

func toolFailedToStart(out []byte) bool { return pyiExtractFailed.Match(out) }

// runParallel executes a contiguous run of Parallel steps with a worker pool
// (MUJARD_BUILD_JOBS, default NumCPU). Outputs are logged in plan order after
// the batch finishes; any failure returns the first failing step's output for
// diagnostics.
func runParallel(ctx context.Context, dir string, steps []Step, log *strings.Builder) (string, error) {
	jobs := runtime.NumCPU()
	if v := os.Getenv("MUJARD_BUILD_JOBS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			jobs = n
		}
	}
	outs := make([]string, len(steps))
	errs := make([]error, len(steps))
	var wg sync.WaitGroup
	sem := make(chan struct{}, jobs)
	for i, s := range steps {
		// Resume support: a finished object is skipped, so a timed-out cold
		// build continues where it stopped; a stale .part means the compile
		// was killed mid-write and is redone.
		if s.Out != "" {
			if st, err := os.Stat(s.Out); err == nil && st.Size() > 0 {
				continue
			}
			os.Remove(s.Out + ".part")
		}
		if err := os.MkdirAll(filepath.Dir(firstOutput(s.Args)), 0o755); err != nil {
			return "", err
		}
		wg.Add(1)
		go func(i int, s Step) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			cmd := exec.CommandContext(ctx, s.Exe, s.Args...)
			cmd.Dir = dir
			out, err := cmd.CombinedOutput()
			outs[i], errs[i] = string(out), err
			if err == nil && s.Out != "" {
				if rerr := os.Rename(s.Out+".part", s.Out); rerr != nil {
					// A parallel job may have landed the same object first.
					if st, serr := os.Stat(s.Out); serr != nil || st.Size() == 0 {
						errs[i] = rerr
					}
					os.Remove(s.Out + ".part")
				}
			}
		}(i, s)
	}
	wg.Wait()
	for i, s := range steps {
		fmt.Fprintf(log, "=== %s\n%s %s\n", s.Name, s.Exe, strings.Join(s.Args, " "))
		log.WriteString(outs[i])
	}
	for _, err := range errs {
		if err != nil {
			return strings.Join(outs, "\n"), err
		}
	}
	return strings.Join(outs, "\n"), ctx.Err()
}

// firstOutput finds the -o target in a step's args so its parent dir can be created.
func firstOutput(args []string) string {
	for i, a := range args {
		if a == "-o" && i+1 < len(args) {
			return args[i+1]
		}
	}
	return "."
}

// parseSize reads `size -A` output and sums the sections a family counts as
// flash vs RAM — the same section sets the platform recipes use for their
// percentage reports.
func parseSize(out, family string) (flash, ram int) {
	var flashSecs, ramSecs map[string]bool
	switch family {
	case "esp32":
		flashSecs = map[string]bool{".iram0.text": true, ".iram0.vectors": true, ".dram0.data": true,
			".dram1.data": true, ".flash.text": true, ".flash.rodata": true, ".flash.appdesc": true,
			".flash.init_array": true, ".eh_frame": true}
		ramSecs = map[string]bool{".dram0.data": true, ".dram0.bss": true, ".dram1.data": true,
			".dram1.bss": true, ".noinit": true}
	case "esp8266": // the core's recipe.size.regex / regex.data
		flashSecs = map[string]bool{".irom0.text": true, ".text": true, ".text1": true, ".data": true, ".rodata": true}
		ramSecs = map[string]bool{".data": true, ".rodata": true, ".bss": true}
	case "stm32": // stm32duino's recipe.size.regex / regex.data
		flashSecs = map[string]bool{".text": true, ".data": true, ".rodata": true}
		ramSecs = map[string]bool{".data": true, ".bss": true, ".noinit": true}
	case "rp2": // arduino-pico's recipe.size.regex / regex.data
		flashSecs = map[string]bool{".boot2": true, ".text": true, ".rodata": true, ".ARM.extab": true, ".ARM.exidx": true}
		ramSecs = map[string]bool{".data": true, ".bss": true, ".ram_vector_table": true, ".uninitialized_ram": true,
			".uninitialized_data": true}
	default: // avr
		flashSecs = map[string]bool{".text": true, ".data": true, ".bootloader": true}
		ramSecs = map[string]bool{".data": true, ".bss": true, ".noinit": true}
	}
	for _, line := range strings.Split(out, "\n") {
		f := strings.Fields(line)
		if len(f) < 2 {
			continue
		}
		n, err := strconv.Atoi(f[1])
		if err != nil {
			continue
		}
		if flashSecs[f[0]] {
			flash += n
		}
		if ramSecs[f[0]] {
			ram += n
		}
	}
	return flash, ram
}
