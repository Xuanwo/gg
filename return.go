package gg

import "io"

type ireturn struct {
	items *group
}

func Return(node ...interface{}) *ireturn {
	i := &ireturn{
		items: newGroup("", "", ", "),
	}
	i.items.append(node...)
	return i
}

func (i *ireturn) render(w io.Writer) {
	writeString(w, "return ")
	i.items.render(w)
}
