package gg

import "io"

type iconst struct {
	items *group
}

func Const() *iconst {
	return &iconst{
		items: newGroup("(", ")", "\n"),
	}
}
func (i *iconst) render(w io.Writer) {
	writeString(w, "const ")
	i.items.render(w)
}

func (i *iconst) Field(name, value string) *iconst {
	i.items.append(&ifield{
		name:      name,
		value:     value,
		separator: "=",
	})
	return i
}
