package gg

import "io"

type ivar struct {
	items *group
}

func Var() *ivar {
	i := &ivar{
		items: newGroup("(", ")", "\n"),
	}
	i.items.omitWrapIf = func() bool {
		// We only need to omit wrap while length == 1.
		// NewIf length == 0, we need to keep it, or it will be invalid expr.
		return i.items.length() == 1
	}
	return i
}

func (i *ivar) render(w io.Writer) {
	writeString(w, "var ")
	i.items.render(w)
}

func (i *ivar) AddField(name, value interface{}) *ivar {
	i.items.append(field(name, value, "="))
	return i
}

func (i *ivar) AddTypedField(name, typ, value interface{}) *ivar {
	i.items.append(typedField(name, typ, value, "="))
	return i
}

func (i *ivar) AddDecl(name, value interface{}) *ivar {
	i.items.append(field(name, value, " "))
	return i
}
