package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"mujard/lang"
)

// boardCaps prints how far each capability is proven on a board (plan Phase 16): `mujard boards caps nano`.
// It reads the compiler's own evidence, so it needs no key and no network.
func boardCaps(args []string, format string, w io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(w, "Say which board: mujard boards caps nano")
		return 2
	}
	b := lang.LookupBoard(args[0])
	if b == nil {
		fmt.Fprintf(w, "Mujard doesn't know the board %q.\n", args[0])
		return 2
	}
	ev := lang.CapabilityEvidence()
	type row struct {
		Capability string `json:"capability"`
		What       string `json:"what"`
		Level      string `json:"level"`
		Evidence   string `json:"evidence,omitempty"`
	}
	var rows []row
	for _, cp := range lang.Capabilities {
		lv := lang.CapabilityLevel(b, cp.Name)
		r := row{Capability: cp.Name, What: cp.What, Level: lv}
		switch lv {
		case lang.LevelHardware:
			r.Evidence = ev.Hardware[b.Name][cp.Name]
		case lang.LevelBuild:
			r.Evidence = "built " + ev.Build[b.Name][cp.Name]
		}
		rows = append(rows, r)
	}
	if format == "json" {
		out, _ := json.MarshalIndent(map[string]any{"board": b.Name, "display": b.Display, "capabilities": rows,
			"notInMujardYet": lang.NotInMujardYet}, "", "  ")
		fmt.Fprintln(w, string(out))
		return 0
	}
	fmt.Fprintf(w, "%s (%s) - what Mujard has proven here\n\n", b.Display, b.Name)
	for _, r := range rows {
		line := fmt.Sprintf("  %-11s %-9s %s", r.Capability, r.Level, r.What)
		if r.Evidence != "" {
			line += " - " + r.Evidence
		}
		fmt.Fprintln(w, line)
	}
	fmt.Fprintf(w, "\nlevels: hardware = ran on a real board; build = a real build succeeded; backend = generated, never built\n"+
		"for this board; language = Mujard has it, this board has no backend; none = the board doesn't have it.\n"+
		"Not in Mujard yet: %s.\n", strings.Join(lang.NotInMujardYet, ", "))
	return 0
}
