package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SchemaV1 is the job/result contract version.
const SchemaV1 = "mujard-build.v1"

// Job is one build request. See docs/toolchain/DIRECT_BUILD_SPEC.md.
type Job struct {
	Schema     string            `json:"schema"`
	Board      string            `json:"board"`
	SketchDir  string            `json:"sketchDir"`
	OutDir     string            `json:"outDir"`
	Options    map[string]string `json:"options"`
	Flags      []string          `json:"flags"`
	NoCache    bool              `json:"noCache"`
	TimeoutSec int               `json:"timeoutSec"`
}

func loadJob(path string) (*Job, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read job: %w", err)
	}
	var j Job
	if err := json.Unmarshal(data, &j); err != nil {
		return nil, fmt.Errorf("bad job json: %w", err)
	}
	if j.Schema != SchemaV1 {
		return nil, fmt.Errorf("bad schema %q (want %s)", j.Schema, SchemaV1)
	}
	if strings.TrimSpace(j.Board) == "" {
		return nil, fmt.Errorf("job has no board")
	}
	if j.SketchDir == "" {
		return nil, fmt.Errorf("job has no sketchDir")
	}
	if j.OutDir == "" {
		return nil, fmt.Errorf("job has no outDir")
	}
	j.SketchDir, err = absClean(j.SketchDir)
	if err != nil {
		return nil, fmt.Errorf("sketchDir: %w", err)
	}
	j.OutDir, err = absClean(j.OutDir)
	if err != nil {
		return nil, fmt.Errorf("outDir: %w", err)
	}
	st, err := os.Stat(j.SketchDir)
	if err != nil || !st.IsDir() {
		return nil, fmt.Errorf("sketchDir %s is not a directory", j.SketchDir)
	}
	if j.TimeoutSec <= 0 {
		j.TimeoutSec = 600
	}
	if j.TimeoutSec > 3600 {
		j.TimeoutSec = 3600
	}
	return &j, nil
}

// absClean makes p absolute and rejects anything that does not resolve cleanly.
func absClean(p string) (string, error) {
	if strings.Contains(p, "\x00") {
		return "", fmt.Errorf("nul byte in path")
	}
	a, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	return filepath.Clean(a), nil
}
