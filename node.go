package gg

import "io"

type Node interface {
	render(w io.Writer)
}

// Embed accept a close clause to build a node.
func Embed(fn func() Node) Node {
	return fn()
}
