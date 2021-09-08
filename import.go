package gg

import "io"

type iimport struct {
	items *group
}

// Import will start a new import group.
func Import() *iimport {
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

// AddPath will import a new path, like `"context"`
func (i *iimport) AddPath(name string) *iimport {
	i.items.append(Lit(name))
	return i
}

// AddDot will import a new path with dot, like `. "context"`
func (i *iimport) AddDot(name string) *iimport {
	i.items.append(String(`. "%s"`, name))
	return i
}

// AddBlank will import a new path with black, like `_ "context"`
func (i *iimport) AddBlank(name string) *iimport {
	i.items.append(String(`_ "%s"`, name))
	return i
}

// AddAlias will import a new path with alias, like `ctx "context"`
func (i *iimport) AddAlias(name, alias string) *iimport {
	i.items.append(String(`%s "%s"`, alias, name))
	return i
}

// AddLine will insert a new line here.
func (i *iimport) AddLine() *iimport {
	i.items.append(Line())
	return i
}

// AddLineComment will insert a new line comment here.
func (i *iimport) AddLineComment(content string, args ...interface{}) *iimport {
	i.items.append(LineComment(content, args...))
	return i
}
