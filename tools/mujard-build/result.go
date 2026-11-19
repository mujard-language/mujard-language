package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Diag is one structured compiler diagnostic: file:line:col: severity: message.
type Diag struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Col      int    `json:"col"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

// Sizes is the firmware/ram report the API and apps display.
type Sizes struct {
	Flash    int `json:"flash"`
	FlashMax int `json:"flashMax"`
	Ram      int `json:"ram"`
	RamMax   int `json:"ramMax"`
}

// Result is the build answer. Callers fall back on Unsupported, report Errors otherwise.
type Result struct {
	Schema      string  `json:"schema"`
	Success     bool    `json:"success"`
	Unsupported bool    `json:"unsupported,omitempty"`
	Board       string  `json:"board,omitempty"`
	Format      string  `json:"format,omitempty"`
	Artifact    string  `json:"artifact,omitempty"`
	Elf         string  `json:"elf,omitempty"`
	Sizes       *Sizes  `json:"sizes,omitempty"`
	Errors      []Diag  `json:"errors"`
	Log         string  `json:"log"`
	Seconds     float64 `json:"seconds"`
}

func writeResult(path string, r *Result) error {
	if r.Errors == nil {
		r.Errors = []Diag{}
	}
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
