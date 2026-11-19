package lang

// The grammar in spec section 3 is normative. These tests read it straight from MUJARD_LANGUAGE_SPEC.md, so the spec
// and the compiler can't drift apart: a recogniser built from the EBNF must accept every good sample, reject every
// parser-error sample, and agree with the hand-written parser on thousands of altered programs.

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var specPath = filepath.Join("..", "..", "..", "docs", "lang", "MUJARD_LANGUAGE_SPEC.md")

// ---- the EBNF, as a tree

type gKind int

const (
	gAlt   gKind = iota // a | b
	gSeq                // a , b
	gOpt                // [ a ]
	gRep                // { a }
	gLook               // & a (PEG lookahead: a must come next, nothing is consumed)
	gNot                // ! a (PEG lookahead: a must not come next)
	gWord               // "x"
	gClass              // NAME, STRING ...
	gRef                // another rule
)

type gNode struct {
	kind gKind
	text string // gWord: the word; gClass, gRef: the name
	kids []*gNode
}

type grammar struct {
	rules map[string]*gNode
	order []string // rule names in the order they are written
}

var tokenClasses = map[string]bool{"NAME": true, "KEYWORD": true, "INT": true, "NUMBER": true, "UNIT": true, "STRING": true,
	"TIMEOFDAY": true, "RAW_BLOCK": true, "NEWLINE": true, "INDENT": true, "DEDENT": true, "EOF": true}

// loadGrammar reads the ```ebnf block of spec section 3.
func loadGrammar(t testing.TB) *grammar {
	t.Helper()
	b, err := os.ReadFile(specPath)
	if err != nil {
		t.Fatal(err)
	}
	spec := strings.ReplaceAll(string(b), "\r\n", "\n")
	start := strings.Index(spec, "```ebnf\n")
	if start < 0 {
		t.Fatal("the spec has no ```ebnf block")
	}
	body := spec[start+len("```ebnf\n"):]
	body = body[:strings.Index(body, "```")]
	g, err := parseEBNF(body)
	if err != nil {
		t.Fatalf("the spec's grammar doesn't parse: %v", err)
	}
	return g
}

var ebnfToken = regexp.MustCompile(`\s*(?:"([^"]*)"|([A-Za-z_][A-Za-z0-9_]*)|([=;|,\[\]{}()&!]))`)

func parseEBNF(src string) (*grammar, error) {
	src = regexp.MustCompile(`(?s)\(\*.*?\*\)`).ReplaceAllString(src, " ")
	type tok struct{ word, name, sym string }
	var toks []tok
	rest := strings.TrimSpace(src)
	for rest != "" {
		m := ebnfToken.FindStringSubmatchIndex(rest)
		if m == nil || m[0] != 0 {
			return nil, fmt.Errorf("can't read %.30q", rest)
		}
		sub := func(i int) string {
			if m[2*i] < 0 {
				return ""
			}
			return rest[m[2*i]:m[2*i+1]]
		}
		t := tok{word: sub(1), name: sub(2), sym: sub(3)}
		if m[2] >= 0 && t.word == "" {
			return nil, fmt.Errorf("empty quoted word")
		}
		toks = append(toks, t)
		rest = strings.TrimSpace(rest[m[1]:])
	}
	i := 0
	peekSym := func(s string) bool { return i < len(toks) && toks[i].sym == s }
	var alt, term func() (*gNode, error)
	term = func() (*gNode, error) {
		if i >= len(toks) {
			return nil, fmt.Errorf("unexpected end")
		}
		t := toks[i]
		i++
		switch {
		case t.word != "":
			return &gNode{kind: gWord, text: t.word}, nil
		case t.name != "" && tokenClasses[t.name]:
			return &gNode{kind: gClass, text: t.name}, nil
		case t.name != "":
			return &gNode{kind: gRef, text: t.name}, nil
		}
		closeFor := map[string]string{"[": "]", "{": "}", "(": ")"}
		if t.sym == "&" || t.sym == "!" {
			inner, err := term()
			if err != nil {
				return nil, err
			}
			if t.sym == "!" {
				return &gNode{kind: gNot, kids: []*gNode{inner}}, nil
			}
			return &gNode{kind: gLook, kids: []*gNode{inner}}, nil
		}
		end, ok := closeFor[t.sym]
		if !ok {
			return nil, fmt.Errorf("unexpected %q", t.sym)
		}
		inner, err := alt()
		if err != nil {
			return nil, err
		}
		if !peekSym(end) {
			return nil, fmt.Errorf("missing %q", end)
		}
		i++
		switch t.sym {
		case "[":
			return &gNode{kind: gOpt, kids: []*gNode{inner}}, nil
		case "{":
			return &gNode{kind: gRep, kids: []*gNode{inner}}, nil
		}
		return inner, nil
	}
	seq := func() (*gNode, error) {
		n := &gNode{kind: gSeq}
		for {
			x, err := term()
			if err != nil {
				return nil, err
			}
			n.kids = append(n.kids, x)
			if !peekSym(",") {
				break
			}
			i++
		}
		if len(n.kids) == 1 {
			return n.kids[0], nil
		}
		return n, nil
	}
	alt = func() (*gNode, error) {
		n := &gNode{kind: gAlt}
		for {
			x, err := seq()
			if err != nil {
				return nil, err
			}
			n.kids = append(n.kids, x)
			if !peekSym("|") {
				break
			}
			i++
		}
		if len(n.kids) == 1 {
			return n.kids[0], nil
		}
		return n, nil
	}
	g := &grammar{rules: map[string]*gNode{}}
	for i < len(toks) {
		name := toks[i].name
		if name == "" || i+1 >= len(toks) || toks[i+1].sym != "=" {
			return nil, fmt.Errorf("expected 'rule =' near token %d", i)
		}
		if _, dup := g.rules[name]; dup {
			return nil, fmt.Errorf("rule %q is defined twice", name)
		}
		i += 2
		body, err := alt()
		if err != nil {
			return nil, fmt.Errorf("in rule %q: %v", name, err)
		}
		if !peekSym(";") {
			return nil, fmt.Errorf("rule %q doesn't end with ';'", name)
		}
		i++
		g.rules[name] = body
		g.order = append(g.order, name)
	}
	return g, nil
}

func (n *gNode) walk(fn func(*gNode)) {
	fn(n)
	for _, k := range n.kids {
		k.walk(fn)
	}
}

// ---- a recogniser: for a rule and a start position, every position a match can end at

type memoKey struct {
	rule *gNode
	pos  int
}

type recogniser struct {
	g       *grammar
	toks    []Token
	memo    map[memoKey][]int
	busy    map[memoKey]bool
	strings map[int]bool // STRING tokens already checked: their {…} parts are expressions
}

func (g *grammar) newRecogniser(toks []Token) *recogniser {
	return &recogniser{g: g, toks: toks, memo: map[memoKey][]int{}, busy: map[memoKey]bool{}, strings: map[int]bool{}}
}

func (g *grammar) accepts(start string, toks []Token) bool {
	return contains(g.newRecogniser(toks).match(&gNode{kind: gRef, text: start}, 0), len(toks))
}

// union merges two sorted, duplicate-free lists of positions.
func union(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	out := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) || j < len(b) {
		switch {
		case j == len(b) || (i < len(a) && a[i] < b[j]):
			out = append(out, a[i])
			i++
		case i == len(a) || b[j] < a[i]:
			out = append(out, b[j])
			j++
		default:
			out = append(out, a[i])
			i++
			j++
		}
	}
	return out
}

func (r *recogniser) match(n *gNode, pos int) []int {
	switch n.kind {
	case gRef:
		rule := r.g.rules[n.text]
		key := memoKey{rule, pos}
		if ends, ok := r.memo[key]; ok {
			return ends
		}
		if r.busy[key] {
			panic("left recursion in rule " + n.text)
		}
		r.busy[key] = true
		ends := r.match(rule, pos)
		delete(r.busy, key)
		r.memo[key] = ends
		return ends
	case gAlt:
		var out []int
		for _, k := range n.kids {
			out = union(out, r.match(k, pos))
		}
		return out
	case gSeq:
		cur := []int{pos}
		for _, k := range n.kids {
			var next []int
			for _, p := range cur {
				next = union(next, r.match(k, p))
			}
			if len(next) == 0 {
				return nil
			}
			cur = next
		}
		return cur
	case gOpt:
		return union([]int{pos}, r.match(n.kids[0], pos))
	case gRep:
		all := []int{pos}
		frontier := []int{pos}
		for len(frontier) > 0 {
			var next []int
			for _, p := range frontier {
				next = union(next, r.match(n.kids[0], p))
			}
			var fresh []int
			for _, e := range next {
				if !contains(all, e) {
					fresh = append(fresh, e)
				}
			}
			all = union(all, fresh)
			frontier = fresh
		}
		return all
	case gLook:
		if len(r.match(n.kids[0], pos)) > 0 {
			return []int{pos}
		}
		return nil
	case gNot:
		if len(r.match(n.kids[0], pos)) == 0 {
			return []int{pos}
		}
		return nil
	case gWord, gClass:
		if pos < len(r.toks) && r.token(n, pos) {
			return []int{pos + 1}
		}
		return nil
	}
	panic("unknown node")
}

func contains(list []int, v int) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

func (r *recogniser) token(n *gNode, pos int) bool {
	t := r.toks[pos]
	if n.kind == gWord {
		return (t.Kind == KEYWORD || t.Kind == OP || t.Kind == NAME || t.Kind == UNIT) && t.Text == n.text
	}
	switch n.text {
	case "NAME":
		return t.Kind == NAME
	case "KEYWORD":
		return t.Kind == KEYWORD
	case "INT":
		return t.Kind == INT
	case "NUMBER":
		return t.Kind == INT || t.Kind == FLOAT
	case "UNIT":
		return t.Kind == UNIT
	case "TIMEOFDAY":
		return t.Kind == TIMEOFDAY
	case "RAW_BLOCK":
		return t.Kind == RAW
	case "NEWLINE":
		return t.Kind == NEWLINE
	case "INDENT":
		return t.Kind == INDENT
	case "DEDENT":
		return t.Kind == DEDENT
	case "EOF":
		return t.Kind == EOF
	case "STRING":
		if t.Kind != STRING {
			return false
		}
		// Every {…} part must be an expression on its own (the parts carry no EOF token).
		if ok, done := r.strings[pos]; done {
			return ok
		}
		ok := true
		for _, part := range t.Parts {
			if part.IsExpr && !r.g.accepts("expr", part.Tokens) {
				ok = false
			}
		}
		r.strings[pos] = ok
		return ok
	}
	return false
}

// ---- tests

func TestGrammarIsComplete(t *testing.T) {
	g := loadGrammar(t)
	used := map[string]bool{"program": true}
	words := map[string]bool{}
	for _, name := range g.order {
		g.rules[name].walk(func(n *gNode) {
			switch n.kind {
			case gRef:
				used[n.text] = true
				if g.rules[n.text] == nil {
					t.Errorf("rule %q uses %q, which isn't defined", name, n.text)
				}
			case gWord:
				words[n.text] = true
			}
		})
	}
	for _, name := range g.order {
		if !used[name] {
			t.Errorf("rule %q is never used", name)
		}
	}
	for kw := range Keywords {
		if !words[kw] {
			t.Errorf("keyword %q (spec 2.4) doesn't appear in the grammar", kw)
		}
	}
	for _, op := range operators {
		if !words[op] {
			t.Errorf("operator %q (spec 2.10) doesn't appear in the grammar", op)
		}
	}
	for w := range words {
		_, isUnit := Units[w]
		isOp := false
		for _, op := range operators {
			isOp = isOp || op == w
		}
		if !Keywords[w] && !isOp && !isUnit && !regexp.MustCompile(`^[a-z_]+$`).MatchString(w) {
			t.Errorf("the grammar uses %q, which is neither a keyword, an operator, a unit nor a context word", w)
		}
	}
}

// mainStart is the grammar rule for a file given to the compiler; a module (no board line) uses 'module'.
func startFor(src string) string {
	f, _ := ParseModuleSource(src)
	if f != nil && f.Board == "" {
		return "module"
	}
	return "main_program"
}

func TestGrammarAcceptsOkSamples(t *testing.T) {
	g := loadGrammar(t)
	for _, s := range loadSamples(t, "ok") {
		toks, errs := Lex(s.src)
		if len(errs) > 0 {
			t.Fatalf("%s: %v", s.name, errs)
		}
		if !g.accepts(startFor(s.src), toks) {
			t.Errorf("%s: the spec's grammar rejects it (%s)", s.name, firstUnparsed(g, startFor(s.src), toks))
		}
	}
}

// grammarOnlyChecks are parser errors for rules the grammar states in a comment rather than in its productions.
var grammarOnlyChecks = map[string]bool{
	"P213": true, // the target of an assignment must be a variable, field or list item
}

func TestGrammarRejectsParserErrorSamples(t *testing.T) {
	g := loadGrammar(t)
	n := 0
	for _, s := range loadSamples(t, "bad") {
		if s.code == "" || s.code[0] != 'P' || grammarOnlyChecks[s.code] {
			continue
		}
		toks, errs := Lex(s.src)
		if len(errs) > 0 {
			continue
		}
		n++
		if g.accepts("main_program", toks) {
			t.Errorf("%s: expected %s, but the spec's grammar accepts it", s.name, s.code)
		}
	}
	if n < 10 {
		t.Errorf("only %d parser-error samples checked", n)
	}
}

// parserAccepts reports whether the hand-written parser finds no error other than the grammar's comment-only rules.
func parserAccepts(src string) (bool, []Diagnostic) {
	_, errs := ParseSource(src)
	for _, e := range errs {
		if !grammarOnlyChecks[e.Code] {
			return false, errs
		}
	}
	return true, errs
}

// mutate makes a small change to a program: the kinds of slips people make, plus random tokens.
func mutate(rng *rand.Rand, src string) string {
	lines := strings.Split(src, "\n")
	pieces := []string{":", ",", "(", ")", "[", "]", "=", "==", "+", "-", ".", "..", "as", "if", "else", "elif", "on",
		"every", "after", "when", "wait", "until", "pin", "analog", "board", "print", "expect", "pass", "return", "x",
		"1", "2 s", "\"t\"", "\"{x}\"", "not", "and", "in", "mod", "fast", "->", "?", "fn", "var", "const", "saved",
		"task", "match", "cpp", "use", "import", "true", "none", "start", "to", "limit", "control", "test", "safe_state"}
	li := rng.Intn(len(lines))
	line := lines[li]
	switch rng.Intn(9) {
	case 0: // delete a line
		lines = append(lines[:li], lines[li+1:]...)
	case 1: // duplicate a line
		lines = append(lines[:li+1], lines[li:]...)
	case 2: // indent a line more
		lines[li] = "    " + line
	case 3: // indent a line less
		lines[li] = strings.TrimPrefix(line, "    ")
	default: // change one word
		words := strings.Fields(line)
		if len(words) == 0 {
			return strings.Join(lines, "\n")
		}
		indent := line[:len(line)-len(strings.TrimLeft(line, " "))]
		wi := rng.Intn(len(words))
		switch rng.Intn(4) {
		case 0:
			words = append(words[:wi], words[wi+1:]...)
		case 1:
			words[wi] = pieces[rng.Intn(len(pieces))]
		case 2:
			words = append(words[:wi], append([]string{pieces[rng.Intn(len(pieces))]}, words[wi:]...)...)
		case 3: // split off a trailing ':' or ',' or ')'
			w := words[wi]
			if len(w) > 1 {
				words[wi] = w[:len(w)-1]
			}
		}
		lines[li] = indent + strings.Join(words, " ")
	}
	return strings.Join(lines, "\n")
}

func TestGrammarAgreesWithParser(t *testing.T) {
	g := loadGrammar(t)
	rng := rand.New(rand.NewSource(20261022))
	samples := loadSamples(t, "ok")
	rounds := 40 // MUJARD_GRAMMAR_ROUNDS=500 for a longer run
	if n, err := strconv.Atoi(os.Getenv("MUJARD_GRAMMAR_ROUNDS")); err == nil && n > 0 {
		rounds = n
	}
	if testing.Short() {
		rounds = 5
	}
	checked, mismatches := 0, 0
	for _, s := range samples {
		if startFor(s.src) != "main_program" {
			continue
		}
		for i := 0; i < rounds; i++ {
			src := s.src
			for k := 0; k <= rng.Intn(2); k++ {
				src = mutate(rng, src)
			}
			toks, lerrs := Lex(src)
			if len(lerrs) > 0 {
				continue
			}
			checked++
			pOK, perrs := parserAccepts(src)
			gOK := g.accepts("main_program", toks)
			if pOK != gOK {
				mismatches++
				if mismatches <= 8 {
					t.Errorf("%s, change %d: parser accepts=%v, grammar accepts=%v; parser says %v\n%s", s.name, i, pOK, gOK,
						perrs, numbered(src))
				}
			}
		}
	}
	if mismatches > 0 {
		t.Errorf("%d of %d altered programs: the parser and the grammar disagree", mismatches, checked)
	}
	t.Logf("%d altered programs: the parser and the grammar agree", checked)
}

func FuzzGrammarAgreesWithParser(f *testing.F) {
	for _, s := range loadSamples(f, "ok") {
		f.Add(s.src)
	}
	g := loadGrammar(f)
	f.Fuzz(func(t *testing.T, src string) {
		toks, lerrs := Lex(src)
		if len(lerrs) > 0 || len(toks) > 400 {
			return
		}
		pOK, perrs := parserAccepts(src)
		if gOK := g.accepts("main_program", toks); pOK != gOK {
			t.Errorf("parser accepts=%v, grammar accepts=%v; parser says %v\n%s", pOK, gOK, perrs, numbered(src))
		}
	})
}

func numbered(src string) string {
	var b strings.Builder
	for i, l := range strings.Split(src, "\n") {
		fmt.Fprintf(&b, "%3d| %s\n", i+1, l)
	}
	return b.String()
}

// firstUnparsed names the furthest token any top-level item reached, to say where the grammar gave up.
func firstUnparsed(g *grammar, start string, toks []Token) string {
	r := g.newRecogniser(toks)
	furthest := 0
	for pos := range toks {
		for _, e := range r.match(&gNode{kind: gRef, text: "top_item"}, pos) {
			if e > furthest {
				furthest = e
			}
		}
	}
	if furthest < len(toks) {
		return "stuck near " + toks[furthest].String()
	}
	return "no complete parse"
}
