// mujard-build is the Mujard-owned build coordinator: it reads one JSON job, resolves the
// board manifest and toolchain package, runs the compiler/linker/image steps itself and
// writes a JSON result. It replaces the arduino-cli orchestration layer for supported
// boards; other boards answer "unsupported" so callers can fall back. See
// docs/toolchain/DIRECT_BUILD_SPEC.md for the contract.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0.0"

func main() {
	jobPath := flag.String("job", "", "JSON job file to build")
	resultPath := flag.String("result", "", "where to write the JSON result (required with -job)")
	planPath := flag.String("plan", "", "print the build steps for this job file, run nothing")
	explainPath := flag.String("explain", "", "explain board/package/vars/sources for this job file")
	list := flag.Bool("list", false, "list every resolvable board name")
	boardsDir := flag.String("boards", "", "board manifest directory (default MUJARD_BOARDS or toolchain/boards)")
	pkgsDir := flag.String("packages", "", "package descriptor directory (default MUJARD_PACKAGES or toolchain/packages)")
	ver := flag.Bool("version", false, "print the version")
	flag.Parse()

	if *ver {
		fmt.Println(version)
		return
	}
	// Subcommand: mujard-build toolchain <dir|list|install|import> — the package
	// manager that installs Mujard-owned toolchain files.
	if flag.NArg() > 0 && flag.Arg(0) == "toolchain" {
		os.Exit(runToolchain(flag.Args()[1:]))
	}
	reg, err := LoadRegistry(findDir(*boardsDir, "MUJARD_BOARDS", "boards"), findDir(*pkgsDir, "MUJARD_PACKAGES", "packages"))
	if err != nil {
		fatalf("mujard-build: %v", err)
	}
	switch {
	case *list:
		for _, line := range reg.List() {
			fmt.Println(line)
		}
	case *planPath != "":
		os.Exit(runJobFile(reg, *planPath, "", modePlan))
	case *explainPath != "":
		os.Exit(runJobFile(reg, *explainPath, "", modeExplain))
	case *jobPath != "":
		if *resultPath == "" {
			*resultPath = filepath.Join(filepath.Dir(*jobPath), "result.json")
		}
		os.Exit(runJobFile(reg, *jobPath, *resultPath, modeBuild))
	default:
		flag.Usage()
		os.Exit(2)
	}
}

func fatalf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(2)
}

// findDir resolves a manifest directory: flag, then env, then the default toolchain/
// subtree found by walking up from the working directory.
func findDir(dir, env, sub string) string {
	if dir != "" {
		return dir
	}
	if v := os.Getenv(env); v != "" {
		return v
	}
	rel := filepath.Join("toolchain", sub)
	cwd, err := os.Getwd()
	if err != nil {
		return rel
	}
	for {
		p := filepath.Join(cwd, rel)
		if st, e := os.Stat(p); e == nil && st.IsDir() {
			return p
		}
		up := filepath.Dir(cwd)
		if up == cwd {
			return rel
		}
		cwd = up
	}
}

const (
	modeBuild = iota
	modePlan
	modeExplain
)

// runJobFile loads the job, resolves it and either builds it or explains it.
func runJobFile(reg *Registry, jobFile, resultPath string, mode int) int {
	job, err := loadJob(jobFile)
	if err != nil {
		if mode == modeBuild {
			writeResult(resultPath, &Result{Schema: SchemaV1, Errors: []Diag{}, Log: "bad job: " + err.Error()})
		}
		fmt.Fprintln(os.Stderr, "mujard-build: "+err.Error())
		return 1
	}
	plan, perr := buildPlan(reg, job)
	switch mode {
	case modePlan:
		if perr != nil {
			fmt.Fprintln(os.Stderr, "mujard-build: "+perr.Error())
			return 1
		}
		for _, s := range plan.Steps {
			fmt.Printf("%s %s\n", s.Exe, strings.Join(s.Args, " "))
		}
		return 0
	case modeExplain:
		if perr != nil {
			fmt.Fprintln(os.Stderr, "mujard-build: "+perr.Error())
			return 1
		}
		fmt.Print(plan.Explain())
		return 0
	}
	res := runBuild(job, plan, perr)
	if werr := writeResult(resultPath, res); werr != nil {
		fmt.Fprintln(os.Stderr, "mujard-build: cannot write result: "+werr.Error())
		return 1
	}
	if !res.Success {
		return 1
	}
	return 0
}
