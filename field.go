package gg

import "io"

type ifield struct {
	name      string
	value     string
	separator string
}

func (f *ifield) render(w io.Writer) {
	writeString(w, f.name, f.separator, f.value)
}

func Field(name, typ string) *ifield {
	return &ifield{
		name:      name,
		value:     typ,
		separator: " ",
	}
}
