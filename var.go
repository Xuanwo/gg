package gg

import "io"

type ivar struct {
	items *group
}

func Var() *ivar {
	return &ivar{
		items: newGroup("(", ")", "\n"),
	}
}

func (i *ivar) render(w io.Writer) {
	writeString(w, "var ")
	i.items.render(w)
}

func (i *ivar) Field(name, value string) *ivar {
	i.items.append(&ifield{
		name:      name,
		value:     value,
		separator: "=",
	})
	return i
}
