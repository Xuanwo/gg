package gg

import "io"

type istruct struct {
	name  string
	items *group
}

func (d *istruct) render(w io.Writer) {
}

func Struct(name string) *istruct {
	return &istruct{
		name:  name,
		items: newGroup("{", "}", "\n"),
	}
}

func (i *istruct) Comment(content string) *istruct {
	return nil
}
func (i *istruct) CommentF(format string, args ...interface{}) *istruct {
	return nil
}

func (i *istruct) Field(name, typ string) *istruct {
	i.items.append(&ifield{
		name:      name,
		value:     typ,
		separator: " ",
	})
	return nil
}
