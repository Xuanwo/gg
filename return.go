package gg

import "io"

type ireturn struct {
	items *group
}

func Return() *ireturn {
	return &ireturn{
		items: newGroup("", "", ", "),
	}
}

func (i *ireturn) render(w io.Writer) {
	writeString(w, "return ")
	i.items.render(w)
}

func (i *ireturn) Id(name string) *ireturn {
	i.items.append(String(name))
	return i
}

func (i *ireturn) Lit(value interface{}) *ireturn {
	i.items.append(Lit(value))
	return i
}

func (i *ireturn) Value(typ string) *ivalue {
	v := Value(typ)
	i.items.append(v)
	return v
}
