package gg

import "io"

type iimport struct {
	items *group
}

func Imports() *iimport {
	return &iimport{
		items: newGroup("(", ")", "\n"),
	}
}

func (i *iimport) render(w io.Writer) {
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
