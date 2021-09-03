package gg

import "io"

type ivalue struct {
	typ   string
	items *group
}

func Value(typ string) *ivalue {
	// FIXME: we need to support builtin type like map and slide.
	return &ivalue{
		typ:   typ,
		items: newGroup("{", "}", ","),
	}
}

func (v *ivalue) render(w io.Writer) {
	writeString(w, v.typ)
	v.items.render(w)
}

func (v *ivalue) Field(name string, value string) *ivalue {
	v.items.append(&ifield{
		name:      name,
		value:     value,
		separator: ":",
	})
	return v
}
