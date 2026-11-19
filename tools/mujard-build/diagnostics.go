package main

import (
	"regexp"
	"strconv"
	"strings"
)

// gccDiag matches the standard GNU diagnostic shapes:
//
//	file.c:12:34: error: message
//	file.c:12: fatal error: message
var (
	gccDiag   = regexp.MustCompile(`^(.+?):(\d+):(\d+):\s*(fatal error|error|warning|note):\s*(.*)$`)
	gccDiagNC = regexp.MustCompile(`^(.+?):(\d+):\s*(fatal error|error|warning|note):\s*(.*)$`)
	ldDiag    = regexp.MustCompile(`^(.+?ld):\s*(.*)$`)
)

// parseDiags turns compiler output into structured errors. #line markers in generated
// sources keep file/line pointed at the .muj file, so no remapping is needed here.
func parseDiags(out string) []Diag {
	var ds []Diag
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimRight(line, "\r")
		if m := gccDiag.FindStringSubmatch(line); m != nil {
			ln, _ := strconv.Atoi(m[2])
			col, _ := strconv.Atoi(m[3])
			ds = append(ds, Diag{File: m[1], Line: ln, Col: col, Severity: m[4], Message: m[5]})
			continue
		}
		if m := gccDiagNC.FindStringSubmatch(line); m != nil {
			ln, _ := strconv.Atoi(m[2])
			ds = append(ds, Diag{File: m[1], Line: ln, Severity: m[3], Message: m[4]})
			continue
		}
		if strings.Contains(line, "error:") {
			ds = append(ds, Diag{Severity: "error", Message: line})
		}
	}
	// errors only — warnings and notes stay in the log, not the error list
	var errs []Diag
	for _, d := range ds {
		if d.Severity == "error" || d.Severity == "fatal error" {
			errs = append(errs, d)
		}
	}
	if len(errs) == 0 && len(ds) > 0 {
		return ds[:1]
	}
	return errs
}
