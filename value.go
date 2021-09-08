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

func (v *ivalue) String() string {
	buf := pool.Get()
	defer buf.Free()

	v.render(buf)
	return buf.String()
}

func (v *ivalue) AddField(name, value interface{}) *ivalue {
	v.items.append(field(name, value, ":"))
	return v
}
