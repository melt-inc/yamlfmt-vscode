package main

import (
	"bufio"
	"strings"
	"syscall/js"

	"github.com/jamesrom/yamlfmt"
)

const DefaultTabSize = 2

func main() {
	js.Global().Set("yamlfmt", js.FuncOf(do))
	<-make(chan bool)
}

func do(_ js.Value, a []js.Value) any {
	args := jsArgs(a)
	if !args.valid() {
		return args[0] // if we can't do anything with the input, return it
	}

	inYaml := args.GetYAML()
	tabSize := args.GetTabSize()
	compact := args.GetCompactSequenceStyle()

	// save the indentation level as it gets stripped, we'll put it back later
	indentation := detectIndentation(inYaml, tabSize)

	formatted, err := yamlfmt.Sfmt(
		inYaml,
		yamlfmt.WithIndentSize(tabSize),
		yamlfmt.WithCompactSequenceStyle(compact),
	)
	if err != nil {
		return args[0]
	}

	out := indent(formatted, indentation)
	return out[:len(out)-1]
}

type jsArgs []js.Value

func (a jsArgs) valid() bool {
	return len(a) == 3 &&
		a[0].Type() == js.TypeString &&
		a[1].Type() == js.TypeObject &&
		a[2].Type() == js.TypeBoolean
}

func (a jsArgs) GetYAML() string {
	return a[0].String()
}

func (a jsArgs) GetTabSize() int {
	ts := a[1].Get("tabSize")
	if ts.Type() != js.TypeNumber {
		return DefaultTabSize
	}
	return ts.Int()
}

func (a jsArgs) GetCompactSequenceStyle() bool {
	return a[2].Truthy()
}

// detectIndentation returns the number of spaces (and tabs, counting tabsize)
// detected at the start of the first line
func detectIndentation(yamlStr string, tabSize int) int {
	count := 0
	for _, char := range yamlStr {
		switch char {
		case '\t':
			count += tabSize
		case ' ':
			count += 1
		default:
			return count
		}
	}
	return count
}

func indent(str string, indentation int) string {
	if indentation == 0 {
		return str
	}
	dent := strings.Repeat(" ", indentation)
	sc := bufio.NewScanner(strings.NewReader(str))
	var b strings.Builder
	b.Grow(len(dent) + len(str)) // lower bound on size
	for sc.Scan() {
		b.WriteString(dent)
		b.WriteString(sc.Text())
		b.WriteRune('\n')
	}
	return b.String()
}
