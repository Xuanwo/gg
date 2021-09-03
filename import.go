package gg

import "io"

type iimport struct {
	items *group
}

func Imports() *iimport {
	i := &iimport{
		items: newGroup("(", ")", "\n"),
	}
	i.items.omitWrapIf = func() bool {
		return i.items.length() <= 1
	}
	return i
}

func (i *iimport) render(w io.Writer) {
	// Don't need to render anything if import is empty
	if i.items.length() == 0 {
		return
	}
	writeString(w, "import ")
	i.items.render(w)
}

func (i *iimport) Path(name string) *iimport {
	i.items.append(Lit(name))
	return i
}
func (i *iimport) Dot(name string) *iimport {
	i.items.append(StringF(`. "%s"`, name))
	return i
}
func (i *iimport) Blank(name string) *iimport {
	i.items.append(StringF(`_ "%s"`, name))
	return i
}
func (i *iimport) Alias(name, alias string) *iimport {
	i.items.append(StringF(`%s "%s"`, alias, name))
	return i
}

func (i *iimport) Line() *iimport {
	i.items.append(Line())
	return i
}
