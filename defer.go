package gg

import "io"

type idefer struct {
	body Node
}

func (i *idefer) render(w io.Writer) {
	writeString(w, "defer ")
	i.body.render(w)
}

func Defer(body interface{}) Node {
	// Add extra space here.
	return &idefer{parseNode(body)}
}
