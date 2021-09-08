package gg

import (
	"fmt"
	"github.com/Xuanwo/go-bufferpool"
	"io"
	"strings"
	"text/template"
)

var pool = bufferpool.New(1024)

// TODO: we will support use to config this logic.
const lineLength = 80

// All internal types are prefixed with `i` to avoid conflict with golang keywords.
type istring string

func (v *istring) render(w io.Writer) {
	writeString(w, string(*v))
}

// S is an alias to String.
// TODO: can we find a new name for this?
var S = String

// String will add a format string in NewGroup, just like fmt.Printf.
func String(format string, args ...interface{}) *istring {
	if len(args) == 0 {
		x := istring(format)
		return &x
	}
	x := istring(fmt.Sprintf(format, args...))
	return &x
}

func formatLineComment(comment string) string {
	buf := pool.Get()
	defer buf.Free()

	// Trim space before going further.
	comment = strings.TrimSpace(comment)

	// Split comment into lines (we will keep original line break.)
	lines := strings.Split(comment, "\n")

	for _, line := range lines {
		cur := 0

		// Start a comment line.
		buf.AppendString("//")

		// Split comment into words
		words := strings.Split(line, " ")

		for _, word := range words {
			// NewIf current line is long enough we need to break it.
			if cur >= lineLength {
				buf.AppendString("\n//")
				cur = 0
			}

			buf.AppendString(" ")
			buf.AppendString(word)
			cur += len(word)
		}
		buf.AppendString("\n")
	}

	return strings.TrimSuffix(buf.String(), "\n")
}

func LineComment(content string, args ...interface{}) *istring {
	if len(args) != 0 {
		content = fmt.Sprintf(content, args...)
	}
	content = formatLineComment(content)
	return String(content)
}

type lit struct {
	value interface{}
}

func Lit(value interface{}) *lit {
	return &lit{value: value}
}

func (v *lit) render(w io.Writer) {
	var out string

	// Code borrowed from github.com/dave/jennifer
	switch v.value.(type) {
	case bool, string, int, complex128:
		out = fmt.Sprintf("%#v", v.value)
	case float64:
		out = fmt.Sprintf("%#v", v.value)
		// NewIf the formatted ivalue is not in scientific notation, and does not have a dot, then
		// we add ".0". Otherwise it will be interpreted as an int.
		// See:
		// https://github.com/dave/jennifer/issues/39
		// https://github.com/golang/go/issues/26363
		if !strings.Contains(out, ".") && !strings.Contains(out, "e") {
			out += ".0"
		}
	case float32, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		out = fmt.Sprintf("%T(%#v)", v.value, v.value)
	case complex64:
		out = fmt.Sprintf("%T%#v", v.value, v.value)
	default:
		panic(fmt.Sprintf("unsupported type for literal: %T", v.value))
	}

	writeString(w, out)
}

func (v *lit) String() string {
	buf := pool.Get()
	defer buf.Free()

	v.render(buf)
	return buf.String()
}

func Line() Node {
	return String("\n")
}

func Template(data interface{}, tmpl string) Node {
	buf := pool.Get()
	defer buf.Free()

	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(buf, data)
	if err != nil {
		panic(fmt.Errorf("template execute: %v", err))
	}
	return String(buf.String())
}
