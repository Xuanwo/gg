package gg

import (
	"io"
)

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
	name      Node
	typ       Node
	value     Node
	separator string
}

func field(name, value interface{}, sep string) *ifield {
	return &ifield{
		name:      parseNode(name),
		value:     parseNode(value),
		separator: sep,
	}
}

func typedField(name, typ, value interface{}, sep string) *ifield {
	return &ifield{
		name:      parseNode(name),
		typ:       parseNode(typ),
		value:     parseNode(value),
		separator: sep,
	}
}

func (f *ifield) render(w io.Writer) {
	f.name.render(w)
	if f.typ != nil {
		writeString(w, " ")
		f.typ.render(w)
	}
	writeString(w, f.separator)
	f.value.render(w)
}
