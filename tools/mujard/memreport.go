package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// The memory a build uses, from the size report arduino-cli prints after linking (its own numbers, not an estimate):
//
//	Sketch uses 4862 bytes (15%) of program storage space. Maximum is 32256 bytes.
//	Global variables use 412 bytes (20%) of dynamic memory, leaving 1636 bytes for local variables. Maximum is 2048 bytes.
var (
	flashUseRx = regexp.MustCompile(`Sketch uses (\d+) bytes .*?Maximum is (\d+) bytes`)
	ramUseRx   = regexp.MustCompile(`Global variables use (\d+) bytes .*?Maximum is (\d+) bytes`)
)

// ramWarnPercent is where a build starts to warn: past it the stack and the program's working values have little
// room left, and on a 2 KB board a crash from running out of RAM looks like random behaviour.
const ramWarnPercent = 75

// buildSizes picks the flash and RAM use out of an arduino-cli log, in the keys the other backends use
// (flash, flashMax, ram, ramMax); nil when the log has no size report.
func buildSizes(log string) map[string]int {
	sizes := map[string]int{}
	if m := flashUseRx.FindStringSubmatch(log); m != nil {
		sizes["flash"], _ = strconv.Atoi(m[1])
		sizes["flashMax"], _ = strconv.Atoi(m[2])
	}
	if m := ramUseRx.FindStringSubmatch(log); m != nil {
		sizes["ram"], _ = strconv.Atoi(m[1])
		sizes["ramMax"], _ = strconv.Atoi(m[2])
	}
	if len(sizes) == 0 {
		return nil
	}
	return sizes
}

func percent(n, of int) int {
	if of <= 0 {
		return 0
	}
	return (n*100 + of/2) / of
}

// tooBig says why a program doesn't fit the board at all ("" when it fits, or when the log had no size report).
func tooBig(sizes map[string]int, board string) string {
	flash, flashMax, ram, ramMax := sizes["flash"], sizes["flashMax"], sizes["ram"], sizes["ramMax"]
	switch {
	case flashMax > 0 && flash > flashMax:
		return fmt.Sprintf("The program is too big for the %s: it needs %d bytes of program memory and the board has %d. "+
			"Remove features or libraries, or use a board with more memory.", board, flash, flashMax)
	case ramMax > 0 && ram > ramMax:
		return fmt.Sprintf("The program's values don't fit in the %s's RAM: they need %d bytes and the board has %d. "+
			"Each text holds 64 characters and each list its full size, so shorten lists or use fewer texts.", board, ram, ramMax)
	}
	return ""
}

// memoryLines are the plain-language lines a build that worked prints about memory: what it uses, and a warning when
// RAM is nearly full.
func memoryLines(sizes map[string]int, board string) []string {
	if sizes == nil {
		return nil
	}
	var out []string
	flash, flashMax, ram, ramMax := sizes["flash"], sizes["flashMax"], sizes["ram"], sizes["ramMax"]
	switch {
	case flashMax > 0 && ramMax > 0:
		out = append(out, fmt.Sprintf("Memory: program %d of %d bytes (%d%%), RAM %d of %d bytes (%d%%).",
			flash, flashMax, percent(flash, flashMax), ram, ramMax, percent(ram, ramMax)))
	case flashMax > 0:
		out = append(out, fmt.Sprintf("Memory: program %d of %d bytes (%d%%).", flash, flashMax, percent(flash, flashMax)))
	}
	if ramMax > 0 && ram <= ramMax && percent(ram, ramMax) > ramWarnPercent {
		out = append(out, fmt.Sprintf("Warning: %d%% of the %s's RAM is used before the program runs, which leaves %d bytes for "+
			"everything it does while running - it may misbehave when that runs out. Shorter lists and fewer texts free RAM.",
			percent(ram, ramMax), board, ramMax-ram))
	}
	return out
}
