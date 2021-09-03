package gg

import "io"

type istruct struct {
	name  string
	items *group
}

func (d *istruct) render(w io.Writer) {
	writeStringF(w, "type %s struct", d.name)
	d.items.render(w)
}

func Struct(name string) *istruct {
	return &istruct{
		name:  name,
		items: newGroup("{", "}", "\n"),
	}
}

func (i *istruct) Line() *istruct {
	i.items.append(Line())
	return i
}

func (i *istruct) Comment(content string) *istruct {
	i.items.append(Comment(content))
	return i
}

func (i *istruct) CommentF(format string, args ...interface{}) *istruct {
	i.items.append(CommentF(format, args...))
	return i
}

func (i *istruct) Field(name, typ string) *istruct {
	i.items.append(&ifield{
		name:      name,
		value:     typ,
		separator: " ",
	})
	return nil
}
