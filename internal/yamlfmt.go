package main

import (
	"bufio"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/jamesrom/yamlfmt"
)

const DefaultTabSize = 2

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("yamlfmt", js.FuncOf(do))
	<-make(chan bool)
}

func do(_ js.Value, args []js.Value) any {
	str := args[0]
	if str.Type() != js.TypeString {
		return str
	}

	opts := args[1]
	tabSize := getTabSize(opts)

	yamlStr := str.String()
	fmt.Println(yamlStr)

	indentation := detectIndentation(yamlStr, tabSize)

	formatted, err := yamlfmt.Sfmt(yamlStr, yamlfmt.WithIndentSize(tabSize))
	if err != nil {
		fmt.Println(err)
		return str
	}

	return indent(formatted, indentation)
}

func getTabSize(opts js.Value) int {
	ts := opts.Get("tabSize")
	if ts.Type() != js.TypeNumber {
		return DefaultTabSize
	}
	return ts.Int()
}

// returns the number of spaces (and tabs, counting tabsize) detected at the start of the first line
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
	dent := strings.Repeat(" ", indentation)
	sc := bufio.NewScanner(strings.NewReader(str))
	var b strings.Builder
	for sc.Scan() {
		b.WriteString(dent)
		b.WriteString(sc.Text())
		b.WriteRune('\n')
	}
	return b.String()
}
