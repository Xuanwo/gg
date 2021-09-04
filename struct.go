package gg

import "io"

type istruct struct {
	name  string
	items *group
}

// Struct will insert a new struct.
func Struct(name string) *istruct {
	return &istruct{
		name: name,
		// We will insert new line before closing the struct to avoid being affect
		// by line comments.
		items: newGroup("{", "\n}", "\n"),
	}
}

func (i *istruct) render(w io.Writer) {
	writeStringF(w, "type %s struct", i.name)
	i.items.render(w)
}

// Line will insert an empty line.
func (i *istruct) Line() *istruct {
	i.items.append(Line())
	return i
}

// LineComment will insert a new line comment.
func (i *istruct) LineComment(content string, args ...interface{}) *istruct {
	i.items.append(LineComment(content, args...))
	return i
}

func (i *istruct) Field(name, typ interface{}) *istruct {
	i.items.append(field(name, typ, " "))
	return i
}
