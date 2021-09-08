package gg

import "io"

type iconst struct {
	items *group
}

func Const() *iconst {
	i := &iconst{
		items: newGroup("(", ")", "\n"),
	}
	i.items.omitWrapIf = func() bool {
		// We only need to omit wrap while length == 1.
		// NewIf length == 0, we need to keep it, or it will be invalid expr.
		return i.items.length() == 1
	}
	return i
}
func (i *iconst) render(w io.Writer) {
	writeString(w, "const ")
	i.items.render(w)
}

func (i *iconst) AddField(name, value interface{}) *iconst {
	i.items.append(field(name, value, "="))
	return i
}
func (i *iconst) AddTypedField(name, typ, value interface{}) *iconst {
	i.items.append(typedField(name, typ, value, "="))
	return i
}

func (i *iconst) AddLineComment(value interface{}) *iconst {
	i.items.append(value)
	return i
}
