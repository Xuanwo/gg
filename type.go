package gg

import "io"

type itype struct {
	name string
	item Node
	sep  string
}

func Type(name string, typ interface{}) *itype {
	return &itype{
		name: name,
		item: parseNode(typ),
	}
}

func TypeAlias(name string, typ interface{}) *itype {
	return &itype{
		name: name,
		item: parseNode(typ),
		sep:  "=",
	}
}

func (i *itype) render(w io.Writer) {
	writeStringF(w, "type %s %s", i.name, i.sep)
	i.item.render(w)
}
