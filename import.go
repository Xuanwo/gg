package gg

import "io"

type iimport struct {
	items *group
}

// Imports will start a new import group.
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

// Path will import a new path, like `"context"`
func (i *iimport) Path(name string) *iimport {
	i.items.append(Lit(name))
	return i
}

// Dot will import a new path with dot, like `. "context"`
func (i *iimport) Dot(name string) *iimport {
	i.items.append(String(`. "%s"`, name))
	return i
}

// Blank will import a new path with black, like `_ "context"`
func (i *iimport) Blank(name string) *iimport {
	i.items.append(String(`_ "%s"`, name))
	return i
}

// Alias will import a new path with alias, like `ctx "context"`
func (i *iimport) Alias(name, alias string) *iimport {
	i.items.append(String(`%s "%s"`, alias, name))
	return i
}

// Line will insert a new line here.
func (i *iimport) Line() *iimport {
	i.items.append(Line())
	return i
}
