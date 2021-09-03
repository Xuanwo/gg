package gg

import "io"

type ipackage struct {
	name string
}

func (i *ipackage) render(w io.Writer) {
	writeStringF(w, "package %s\n", i.name)
}

func Package(name string) *ipackage {
	return &ipackage{name: name}
}
