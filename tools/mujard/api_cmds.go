package main

// The small commands over the developer API. Each is a thin wrapper: parse, call, print.
// Read commands take --format json, which prints the server's exact JSON document.
import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// showAnswer prints a decoded API answer: raw with --format json, else indented JSON.
func showAnswer(w io.Writer, raw []byte, asJSON bool) {
	if asJSON {
		fmt.Fprintln(w, string(raw))
		return
	}
	var v any
	if json.Unmarshal(raw, &v) == nil {
		b, _ := json.MarshalIndent(v, "", "  ")
		fmt.Fprintln(w, string(b))
		return
	}
	fmt.Fprintln(w, string(raw))
}

func runMe(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("me", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	format := fs.String("format", "", "\"json\" prints the raw document")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	cl.finish()
	if cl.needKey(w) {
		return 2
	}
	status, raw, _, err := cl.callRaw("GET", "/me", nil, nil)
	if err != nil {
		fmt.Fprintf(w, "Cannot reach the API: %s\n", short(err))
		return 1
	}
	var doc map[string]any
	_ = json.Unmarshal(raw, &doc)
	if status != 200 {
		return failAPI(w, status, doc)
	}
	if *format == "json" {
		showAnswer(w, raw, true)
		return 0
	}
	k, _ := doc["key"].(map[string]any)
	d, _ := doc["developer"].(map[string]any)
	if k != nil {
		fmt.Fprintf(w, "key       %q (%s)  scopes: %s\n", str(k, "name"), str(k, "mode"), strings.Join(strList(k["scopes"]), ", "))
		fmt.Fprintf(w, "expires   %s   rate %d/min   daily quota %d\n", str(k, "expires_at"), num(k, "rate_per_minute"), num(k, "daily_quota"))
	}
	if d != nil {
		fmt.Fprintf(w, "developer %s   spent this month: %v coin(s) (cap %v)\n", str(d, "company"), d["spent_this_month"], d["monthly_spend_cap"])
	}
	return 0
}

func runUsage(args []string, w io.Writer) int {
	// "usage windows" is a different thing: this desktop's own session/week allowance, answered by the desktop app.
	if firstArg(args) == "windows" {
		return runDesktopNew(w, "usage", append([]string{"windows"}, formatArgs(args)...))
	}
	fs := flag.NewFlagSet("usage", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	format := fs.String("format", "", "\"json\" prints the raw document")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	cl.finish()
	if cl.needKey(w) {
		return 2
	}
	status, raw, _, err := cl.callRaw("GET", "/usage", nil, nil)
	if err != nil {
		fmt.Fprintf(w, "Cannot reach the API: %s\n", short(err))
		return 1
	}
	if status != 200 {
		var doc map[string]any
		_ = json.Unmarshal(raw, &doc)
		return failAPI(w, status, doc)
	}
	showAnswer(w, raw, *format == "json")
	return 0
}

func runBoards(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("boards", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	format := fs.String("format", "", "\"json\" prints the raw document")
	onlyAvailable := fs.Bool("available", false, "search: only boards that can be built right now")
	desktop := fs.Bool("desktop", false, "show the desktop app's own board catalog (protocol, baud, USB ids, flashable) instead of the build API's")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	if *desktop {
		return desktopBoards(rest, *format, w)
	}
	if firstArg(rest) == "caps" { // the compiler's own capability table: local, no key
		return boardCaps(rest[1:], *format, w)
	}
	// "boards list" is the mujard-style spelling of plain "boards".
	if firstArg(rest) == "list" {
		rest = rest[1:]
		if len(rest) > 0 {
			fmt.Fprintln(w, "\"mujard boards list\" takes nothing after \"list\".")
			return 2
		}
	}
	// "boards info <id>" is "boards <id>" (the noun-first spelling of "board info <id>").
	if len(rest) > 0 && rest[0] == "info" {
		if len(rest) != 2 {
			fmt.Fprintln(w, "Say which board: mujard boards info nano   (\"mujard boards\" lists them)")
			return 2
		}
		rest = rest[1:]
	}
	search := ""
	if len(rest) > 0 && rest[0] == "search" {
		search = strings.ToLower(strings.TrimSpace(strings.Join(rest[1:], " ")))
		if search == "" {
			fmt.Fprintln(w, "Search for what? mujard boards search esp32   (matches the id, name and family)")
			return 2
		}
		rest = nil
	}
	cl.finish()
	if cl.needKey(w) {
		return 2
	}
	path := "/boards"
	if id := firstArg(rest); id != "" {
		path += "/" + id
	}
	status, raw, _, err := cl.callRaw("GET", path, nil, nil)
	if err != nil {
		fmt.Fprintf(w, "Cannot reach the API: %s\n", short(err))
		return 1
	}
	var doc map[string]any
	_ = json.Unmarshal(raw, &doc)
	if status != 200 {
		return failAPI(w, status, doc)
	}
	if search != "" {
		all, _ := doc["data"].([]any)
		matches := []any{}
		for _, e := range all {
			b, _ := e.(map[string]any)
			if b == nil || (*onlyAvailable && b["available"] != true) {
				continue
			}
			if strings.Contains(strings.ToLower(str(b, "id")+" "+str(b, "name")+" "+str(b, "family")), search) {
				matches = append(matches, e)
			}
		}
		if *format == "json" {
			out, _ := json.Marshal(map[string]any{"data": matches})
			fmt.Fprintln(w, string(out))
			return 0
		}
		if len(matches) == 0 {
			fmt.Fprintf(w, "No board matches %q. \"mujard boards\" lists them all.\n", search)
			return 0
		}
		printBoardTable(w, matches)
		fmt.Fprintf(w, "\n%d match(es) for %q.\n", len(matches), search)
		return 0
	}
	if *format == "json" || firstArg(rest) != "" {
		showAnswer(w, raw, *format == "json")
		return 0
	}
	list, _ := doc["data"].([]any)
	printBoardTable(w, list)
	fmt.Fprintf(w, "\n%d boards. \"mujard boards <id>\" for one; build with: mujard build <id> <sketch>\n", len(list))
	return 0
}

func printBoardTable(w io.Writer, list []any) {
	fmt.Fprintf(w, "%-22s %-38s %-16s %s\n", "id", "name", "family", "buildable now")
	for _, e := range list {
		b, _ := e.(map[string]any)
		if b == nil {
			continue
		}
		av := "yes"
		if b["available"] != true {
			av = "no (its package is not on the build server)"
		}
		fmt.Fprintf(w, "%-22s %-38s %-16s %s\n", str(b, "id"), str(b, "name"), str(b, "family"), av)
	}
}

// "board" is the mujard-style spelling: "mujard board info <id>" (or just "mujard board <id>") asks for one
// board's detail - the same call "mujard boards <id>" makes.
func runBoardInfo(args []string, w io.Writer) int {
	rest := args
	if len(rest) > 0 && rest[0] == "info" {
		rest = rest[1:]
	}
	fs := flag.NewFlagSet("board", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	addAPIFlags(fs)
	fs.String("format", "", "")
	fs.Bool("desktop", false, "")
	positional, err := parseArgs(fs, rest)
	if err != nil || firstArg(positional) == "" {
		fmt.Fprintln(w, "Say which board: mujard board info nano   (\"mujard boards\" lists them)")
		return 2
	}
	return runBoards(rest, w)
}

// isModelCommand: "mujard ai model list" / "mujard ai model info <id>" are read-only catalog views the
// desktop app answers from models.php. A bare "mujard ai model" or "mujard ai model use ..." goes there
// too, for its usage error - never to the paid AI as a prompt. Every other "mujard ai ..." is a
// prompt, exactly as before.
func isModelCommand(args []string) bool {
	if len(args) == 0 || args[0] != "model" {
		return false
	}
	return len(args) == 1 || args[1] == "list" || args[1] == "info" || args[1] == "use"
}

func runAI(args []string, w io.Writer) int {
	if isModelCommand(args) {
		return runDesktop(w, "ai", args)
	}
	// --mode / --file / --model: a question for the AI with files and a mode, asked through the desktop app
	// (the developer API's chat takes only a prompt). It spends THIS DESKTOP's allowance, not the developer key's coins.
	if isDesktopAIRequest(args) {
		return runDesktopAI(args, w)
	}
	fs := flag.NewFlagSet("ai", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	cl.finish()
	prompt := strings.Join(rest, " ")
	if strings.TrimSpace(prompt) == "" {
		fmt.Fprintln(w, "Usage: mujard ai \"your question\"")
		return 2
	}
	if cl.needKey(w) {
		return 2
	}
	status, doc, err := cl.call("POST", "/ai/chat",
		map[string]any{"prompt": prompt},
		map[string]string{"Idempotency-Key": idemKey()})
	if err != nil {
		fmt.Fprintf(w, "Cannot reach the API: %s\n", short(err))
		return 1
	}
	if status != 200 && status != 202 {
		return failAPI(w, status, doc)
	}
	switch str(doc, "state") {
	case "done":
		fmt.Fprintln(w, str(doc, "answer"))
		if c := num(doc, "cost"); c > 0 {
			fmt.Fprintf(w, "(%d coin(s))\n", c)
		}
	default:
		fmt.Fprintf(w, "The AI call is %s (id %s).\n", str(doc, "state"), str(doc, "id"))
		if e := str(doc, "error"); e != "" {
			fmt.Fprintln(w, e)
		}
	}
	return 0
}

var webhookIDRx = regexp.MustCompile(`^\d{1,9}$`)

// runWebhooks: mujard webhooks list | add <url> [--events a,b] | del <id> | test <id> | deliveries <id>
func runWebhooks(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("webhooks", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	events := fs.String("events", "build.completed,build.failed", "comma-separated events for \"add\"")
	yes := fs.Bool("yes", false, "rotate-secret: do not ask first")
	rest, perr := parseArgs(fs, args)
	if perr != nil {
		return 2
	}
	cl.finish()
	if cl.needKey(w) {
		return 2
	}
	sub := ""
	if len(rest) > 0 {
		sub = rest[0]
	}
	var status int
	var raw []byte
	var err error
	switch sub {
	case "list", "":
		status, raw, _, err = cl.callRaw("GET", "/webhooks", nil, nil)
	case "add":
		if len(rest) < 2 {
			fmt.Fprintln(w, "Usage: mujard webhooks add <https-url> [--events build.completed,build.failed]")
			return 2
		}
		evs := []string{}
		for _, e := range strings.Split(*events, ",") {
			if s := strings.TrimSpace(e); s != "" {
				evs = append(evs, s)
			}
		}
		status, raw, _, err = cl.callRaw("POST", "/webhooks",
			map[string]any{"url": rest[1], "events": evs},
			map[string]string{"Idempotency-Key": idemKey()})
	case "del":
		if len(rest) < 2 {
			fmt.Fprintln(w, "Usage: mujard webhooks del <id>")
			return 2
		}
		status, raw, _, err = cl.callRaw("DELETE", "/webhooks/"+rest[1], nil, nil)
	case "test":
		if len(rest) < 2 {
			fmt.Fprintln(w, "Usage: mujard webhooks test <id>")
			return 2
		}
		status, raw, _, err = cl.callRaw("POST", "/webhooks/"+rest[1]+"/test", nil,
			map[string]string{"Idempotency-Key": idemKey()})
	case "deliveries":
		if len(rest) < 2 {
			fmt.Fprintln(w, "Usage: mujard webhooks deliveries <id>")
			return 2
		}
		status, raw, _, err = cl.callRaw("GET", "/webhooks/"+rest[1]+"/deliveries", nil, nil)
	case "info", "get", "enable", "rotate-secret":
		if len(rest) != 2 || !webhookIDRx.MatchString(rest[1]) {
			fmt.Fprintf(w, "Usage: mujard webhooks %s <id>   (\"mujard webhooks list\" shows the ids)\n", sub)
			return 2
		}
		switch sub {
		case "info", "get":
			status, raw, _, err = cl.callRaw("GET", "/webhooks/"+rest[1], nil, nil)
		case "enable":
			status, raw, _, err = cl.callRaw("POST", "/webhooks/"+rest[1]+"/enable", nil,
				map[string]string{"Idempotency-Key": idemKey()})
		default:
			if !confirmed(w, *yes, "Make a new signing secret for webhook "+rest[1]+"? The old secret stops working at once.") {
				fmt.Fprintln(w, "Not changed.")
				return 1
			}
			status, raw, _, err = cl.callRaw("POST", "/webhooks/"+rest[1]+"/rotate-secret", nil,
				map[string]string{"Idempotency-Key": idemKey()})
		}
	default:
		fmt.Fprintf(w, "unknown webhooks command %q: list | add | del | test | deliveries | info | enable | rotate-secret\n", sub)
		return 2
	}
	if err != nil {
		fmt.Fprintf(w, "Cannot reach the API: %s\n", short(err))
		return 1
	}
	var doc map[string]any
	_ = json.Unmarshal(raw, &doc)
	if status < 200 || status > 299 {
		return failAPI(w, status, doc)
	}
	if sub == "add" {
		// The signing secret is shown once by the API and never again - keep it visible here.
		fmt.Fprintln(w, "Webhook created. Keep its signing secret safe:")
	}
	if sub == "rotate-secret" {
		fmt.Fprintln(w, "New signing secret - it is shown only this once. The old one no longer signs messages:")
	}
	showAnswer(w, raw, false)
	return 0
}

func strList(v any) []string {
	out := []string{}
	if l, _ := v.([]any); l != nil {
		for _, e := range l {
			if s, ok := e.(string); ok {
				out = append(out, s)
			}
		}
	}
	sort.Strings(out)
	return out
}
