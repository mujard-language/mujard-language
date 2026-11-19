package lang

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Recursion (plan Phase 15, "stronger memory diagnostics"): a function that calls itself, directly or through others,
// takes more stack RAM with every call, and on an AVR board's 2 KB a deep recursion silently overwrites the program's
// variables. The RAM estimate (W614) can't see it, so the call graph of the program's own functions is checked: each
// cycle is W617 at the first function in it (an error under 'safety strict', spec 6.11).

// walkTree calls visit for n and every node inside it, by the syntax tree's own fields (so a new kind of node is
// walked without changes here, unlike codegen's eachNode).
func walkTree(n any, visit func(Node)) {
	var walk func(v reflect.Value)
	nodeType := reflect.TypeOf((*Node)(nil)).Elem()
	walk = func(v reflect.Value) {
		switch v.Kind() {
		case reflect.Interface, reflect.Pointer:
			if v.IsNil() {
				return
			}
			if v.Type().Implements(nodeType) {
				if node, ok := v.Interface().(Node); ok {
					visit(node)
				}
			}
			walk(v.Elem())
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				if v.Type().Field(i).IsExported() {
					walk(v.Field(i))
				}
			}
		case reflect.Slice, reflect.Array:
			for i := 0; i < v.Len(); i++ {
				walk(v.Index(i))
			}
		}
	}
	walk(reflect.ValueOf(n))
}

// recursionCheck reports W617 for every cycle in the calls between the program's own functions.
func (c *checker) recursionCheck() {
	fns := map[string]*FnDecl{}
	for _, it := range c.file.Items {
		if fd, ok := it.(*FnDecl); ok {
			fns[fd.Name] = fd
		}
	}
	calls := map[string][]string{}
	for name, fd := range fns {
		seen := map[string]bool{}
		walkTree(fd.Body, func(n Node) {
			if call, ok := n.(*Call); ok {
				if id, ok := call.Fn.(*Ident); ok && fns[id.Name] != nil && !seen[id.Name] {
					seen[id.Name] = true
					calls[name] = append(calls[name], id.Name)
				}
			}
		})
		sort.Strings(calls[name])
	}
	names := make([]string, 0, len(fns))
	for n := range fns {
		names = append(names, n)
	}
	sort.Slice(names, func(i, j int) bool { // report in source order
		a, b := fns[names[i]].At, fns[names[j]].At
		return a.Line < b.Line || a.Line == b.Line && a.Col < b.Col
	})
	reported := map[string]bool{}
	for _, start := range names {
		if reported[start] {
			continue
		}
		// the shortest way back to start, if there is one (breadth first, so the chain shown is the plainest)
		prev := map[string]string{}
		queue := []string{start}
		found := false
		for len(queue) > 0 && !found {
			at := queue[0]
			queue = queue[1:]
			for _, next := range calls[at] {
				if next == start {
					prev[start] = at
					found = true
					break
				}
				if _, ok := prev[next]; !ok && next != start {
					prev[next] = at
					queue = append(queue, next)
				}
			}
		}
		if !found {
			continue
		}
		var path []string // the functions start goes through, in call order
		for at := prev[start]; at != start; at = prev[at] {
			path = append([]string{at}, path...)
		}
		chain := append([]string{start}, path...)
		for _, n := range chain {
			reported[n] = true // one warning per cycle
		}
		through := ""
		if len(chain) > 1 {
			through = " (through " + strings.Join(chain[1:], ", ") + ")"
		}
		board, ram := "board", "a few"
		if c.board != nil {
			board = c.board.Display
			if c.board.Memory.RAM > 0 {
				ram = strconv.Itoa(c.board.Memory.RAM)
			}
		}
		c.errorf("W617", fns[start].At, "fn", start, "through", through, "board", board, "ram", ram)
	}
}
