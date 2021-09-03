package gg

import "io"

// ifield is used to represent a key-value pair.
//
// It could be used in:
// - struct decl
// - struct value
// - method receiver
// - function parameter
// - function result
// - ...
type ifield struct {
	name      string
	value     string
	separator string
}

func (f *ifield) render(w io.Writer) {
	writeString(w, f.name, f.separator, f.value)
}
