package gg

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

type istring string

func (v *istring) render(w io.Writer) {
	writeString(w, string(*v))
}

func String(v string) *istring {
	x := istring(v)
	return &x
}

func StringF(format string, args ...interface{}) *istring {
	x := istring(fmt.Sprintf(format, args...))
	return &x
}

func Comment(content string) *istring {
	content = formatComment(content)
	return String(content)
}

func CommentF(content string, args ...interface{}) *istring {
	content = fmt.Sprintf(content, args...)
	content = formatComment(content)
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
		// If the formatted ivalue is not in scientific notation, and does not have a dot, then
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
