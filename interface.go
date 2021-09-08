package gg

import "io"

type isignature struct {
	comments   *group
	name       string
	parameters *group
	results    *group
}

func signature(name string) *isignature {
	i := &isignature{
		name:       name,
		comments:   newGroup("", "", "\n"),
		parameters: newGroup("(", ")", ","),
		results:    newGroup("(", ")", ","),
	}
	// We should omit the `()` if result is empty
	// Read about omit in NewFunction comments.
	i.results.omitWrapIf = func() bool {
		l := i.results.length()
		if l == 0 {
			// There is no result fields, we can omit `()` safely.
			return true
		}
		return false
	}
	return i
}

func (i *isignature) render(w io.Writer) {
	// Render function name
	writeString(w, i.name)

	// Render parameters
	i.parameters.render(w)
	// Render results
	i.results.render(w)
}

func (i *isignature) AddParameter(name, typ interface{}) *isignature {
	i.parameters.append(field(name, typ, " "))
	return i
}

func (i *isignature) AddResult(name, typ interface{}) *isignature {
	i.results.append(field(name, typ, " "))
	return i
}

type iinterface struct {
	name  string
	items *group
}

func Interface(name string) *iinterface {
	return &iinterface{
		name:  name,
		items: newGroup("{\n", "}", "\n"),
	}
}

func (i *iinterface) render(w io.Writer) {
	writeStringF(w, "type %s interface", i.name)
	i.items.render(w)
}

func (i *iinterface) NewFunction(name string) *isignature {
	sig := signature(name)
	i.items.append(sig)
	return sig
}

// AddLineComment will insert a new line comment.
func (i *iinterface) AddLineComment(content string, args ...interface{}) *iinterface {
	i.items.append(LineComment(content, args...))
	return i
}

// AddLine will insert a new line.
func (i *iinterface) AddLine() *iinterface {
	i.items.append(Line())
	return i
}
