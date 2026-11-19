package main

// mujard builds list | get | artifact | cancel - the developer API's build routes that "mujard build" only
// uses internally: GET /v1/builds, GET /v1/builds/{id}, GET /v1/builds/{id}/artifact and
// DELETE /v1/builds/{id}. Listing and reading are free; cancel is refunded by the server.
// "mujard build get|status|cancel|artifact bld_..." is accepted too, but only when the id really is a
// build id, so a board that happened to be called "get" could never be mistaken for the sub-command.
import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var buildIDRx = regexp.MustCompile(`^bld_[0-9a-f]{16}$`)

// stdinReader is where a confirm question is answered from; tests replace it.
var stdinReader io.Reader = os.Stdin

// confirmed asks a yes/no question unless --yes was given. Anything but "y"/"yes" - including the
// end-of-input a script or pipe gives - counts as no, so nothing is changed unprompted.
func confirmed(w io.Writer, yes bool, question string) bool {
	if yes {
		return true
	}
	fmt.Fprintf(w, "%s [y/N] ", question)
	line, _ := bufio.NewReader(stdinReader).ReadString('\n')
	a := strings.ToLower(strings.TrimSpace(line))
	return a == "y" || a == "yes"
}

func buildSubAlias(word string) (string, bool) {
	switch word {
	case "get", "status":
		return "get", true
	case "artifact", "cancel":
		return word, true
	}
	return "", false
}

func runBuilds(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("builds", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	format := fs.String("format", "", "\"json\" prints the raw document")
	state := fs.String("state", "", "list: only builds in this state (queued, running, done, failed, cancelled)")
	mode := fs.String("mode", "", "list: only test or live builds")
	limit := fs.Int("limit", 20, "list: how many builds (1-100)")
	after := fs.String("after", "", "list: continue after this build id (the previous page's last one)")
	out := fs.String("out", "", "artifact: where the firmware goes (default: <build id>.<format> here)")
	force := fs.Bool("force", false, "artifact: replace the file if it exists")
	yes := fs.Bool("yes", false, "cancel: do not ask")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	sub := "list"
	if len(rest) > 0 {
		sub = rest[0]
		rest = rest[1:]
	}
	asJSON := *format == "json"
	if *format != "" && !asJSON {
		fmt.Fprintln(w, "--format is \"json\" or nothing.")
		return 2
	}
	idArg := func() (string, bool) {
		if len(rest) != 1 {
			fmt.Fprintf(w, "Usage: mujard builds %s <build id>   (ids look like bld_0123456789abcdef; \"mujard builds list\" shows them)\n", sub)
			return "", false
		}
		if !buildIDRx.MatchString(rest[0]) {
			fmt.Fprintf(w, "%q is not a build id (they look like bld_0123456789abcdef).\n", rest[0])
			return "", false
		}
		return rest[0], true
	}
	switch sub {
	case "list":
		if len(rest) > 0 {
			fmt.Fprintln(w, "\"mujard builds list\" takes options only: --state --mode --limit --after")
			return 2
		}
	case "get", "artifact", "cancel":
		if _, ok := idArg(); !ok {
			return 2
		}
	default:
		fmt.Fprintf(w, "unknown builds command %q: list | get <id> | artifact <id> | cancel <id>\n", sub)
		return 2
	}
	cl.finish()
	if cl.needKey(w) {
		return 2
	}

	switch sub {
	case "list":
		q := url.Values{}
		if *limit < 1 || *limit > 100 {
			fmt.Fprintln(w, "--limit is a whole number from 1 to 100.")
			return 2
		}
		q.Set("limit", strconv.Itoa(*limit))
		if *state != "" {
			switch *state {
			case "queued", "running", "done", "failed", "cancelled":
				q.Set("state", *state)
			default:
				fmt.Fprintln(w, "--state is one of queued, running, done, failed, cancelled.")
				return 2
			}
		}
		if *mode != "" {
			if *mode != "test" && *mode != "live" {
				fmt.Fprintln(w, "--mode is test or live.")
				return 2
			}
			q.Set("mode", *mode)
		}
		if *after != "" {
			if !buildIDRx.MatchString(*after) {
				fmt.Fprintln(w, "--after is a build id from a previous page.")
				return 2
			}
			q.Set("starting_after", *after)
		}
		status, raw, _, cerr := cl.callRaw("GET", "/builds?"+q.Encode(), nil, nil)
		doc, code := apiResult(w, status, raw, cerr)
		if code != 0 {
			return code
		}
		if asJSON {
			showAnswer(w, raw, true)
			return 0
		}
		list, _ := doc["data"].([]any)
		if len(list) == 0 {
			fmt.Fprintln(w, "No builds match.")
			return 0
		}
		fmt.Fprintf(w, "%-22s %-14s %-10s %-5s %-5s %s\n", "id", "board", "state", "mode", "cost", "created")
		for _, e := range list {
			b, _ := e.(map[string]any)
			if b == nil {
				continue
			}
			fmt.Fprintf(w, "%-22s %-14s %-10s %-5s %-5d %s\n", str(b, "id"), str(b, "board"), str(b, "state"), str(b, "mode"), num(b, "cost"), str(b, "created_at"))
		}
		if doc["has_more"] == true {
			fmt.Fprintf(w, "\nMore builds: mujard builds list --after %s\n", str(doc, "next_cursor"))
		}
		return 0

	case "get":
		status, raw, _, cerr := cl.callRaw("GET", "/builds/"+rest[0], nil, nil)
		doc, code := apiResult(w, status, raw, cerr)
		if code != 0 {
			return code
		}
		if asJSON {
			showAnswer(w, raw, true)
			return 0
		}
		printBuildSummary(w, doc)
		return 0

	case "artifact":
		id := rest[0]
		status, raw, _, cerr := cl.callRaw("GET", "/builds/"+id, nil, nil)
		job, code := apiResult(w, status, raw, cerr)
		if code != 0 {
			return code
		}
		result, _ := job["result"].(map[string]any)
		if str(job, "state") != "done" || job["has_artifact"] != true {
			fmt.Fprintf(w, "Build %s is %s and has no firmware to download", id, str(job, "state"))
			if str(job, "mode") == "test" {
				fmt.Fprint(w, " (a test key builds nothing for real)")
			}
			fmt.Fprintln(w, ".")
			return 1
		}
		ext := "hex"
		if f := str(result, "format"); f == "bin" || f == "uf2" {
			ext = f
		}
		path := *out
		if path == "" {
			path = id + "." + ext
		}
		if _, serr := os.Stat(path); serr == nil && !*force {
			fmt.Fprintf(w, "%s already exists - pass --force to replace it, or --out to pick another file.\n", path)
			return 1
		}
		status, body, _, cerr := cl.callRaw("GET", "/builds/"+id+"/artifact", nil, nil)
		if cerr != nil {
			fmt.Fprintf(w, "Cannot reach the API: %s\n", short(cerr))
			return 1
		}
		if status != 200 {
			var d map[string]any
			_ = json.Unmarshal(body, &d)
			return failAPI(w, status, d)
		}
		sum := sha256.Sum256(body)
		if werr := os.WriteFile(filepath.Clean(path), body, 0o644); werr != nil {
			fmt.Fprintf(w, "The firmware could not be saved to %s: %s\n", path, werr)
			return 1
		}
		fmt.Fprintf(w, "Saved %d bytes of %s from build %s -> %s (sha256 %s)\n", len(body), ext, id, path, hex.EncodeToString(sum[:])[:12])
		return 0

	case "cancel":
		if !confirmed(w, *yes, "Cancel build "+rest[0]+"? (only a build that has not started can be cancelled; the server refunds it)") {
			fmt.Fprintln(w, "Not cancelled.")
			return 1
		}
		status, raw, _, cerr := cl.callRaw("DELETE", "/builds/"+rest[0], nil, nil)
		if _, code := apiResult(w, status, raw, cerr); code != 0 {
			return code
		}
		showAnswer(w, raw, asJSON)
		return 0
	}
	return 2
}

// apiResult turns one finished call into the decoded document, or the exit code to return after
// printing why it failed. It is the same wording the older commands use.
func apiResult(w io.Writer, status int, raw []byte, err error) (map[string]any, int) {
	if err != nil {
		fmt.Fprintf(w, "Cannot reach the API: %s\n", short(err))
		return nil, 1
	}
	var doc map[string]any
	_ = json.Unmarshal(raw, &doc)
	if status < 200 || status > 299 {
		return nil, failAPI(w, status, doc)
	}
	return doc, 0
}

func printBuildSummary(w io.Writer, job map[string]any) {
	fmt.Fprintf(w, "build    %s\nboard    %s\nstate    %s (%s key)\ncreated  %s\n", str(job, "id"), str(job, "board"), str(job, "state"), str(job, "mode"), str(job, "created_at"))
	fmt.Fprintf(w, "cost     %d coin(s)%s\n", num(job, "cost"), map[bool]string{true: " - refunded", false: ""}[job["refunded"] == true])
	result, _ := job["result"].(map[string]any)
	if result != nil {
		if result["success"] == true {
			fmt.Fprintf(w, "result   built (%s)\n", str(result, "format"))
		} else {
			fmt.Fprintln(w, "result   did not produce firmware")
		}
		printBuildErrors(w, result)
		printSizes(w, result)
	}
	if job["has_artifact"] == true {
		fmt.Fprintf(w, "\nDownload the firmware: mujard builds artifact %s\n", str(job, "id"))
	}
}
