package gg

import (
	"fmt"
	"io"
)

type Node interface {
	render(w io.Writer)
}

// Embed accept a close clause to build a node.
func Embed(fn func() Node) Node {
	return fn()
}

// parseNodes will parse valid input into node slices.
func parseNodes(in []interface{}) []Node {
	ns := make([]Node, 0, len(in))
	for _, v := range in {
		ns = append(ns, parseNode(v))
	}
	return ns
}

// parseNode will parse a valid input into a node.
// For now, we only support two types:
// - Native Node
// - golang string
func parseNode(in interface{}) Node {
	switch v := in.(type) {
	case Node:
		return v
	case string:
		return String(v)
	default:
		panic(fmt.Errorf("invalid input: %s", v))
	}
}
