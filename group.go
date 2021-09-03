package gg

import (
	"fmt"
	"io"
	"os"
)

func Group() *group {
	return newGroup("", "", "\n")
}

func newGroup(open, close, sep string) *group {
	return &group{
		open:      open,
		close:     close,
		separator: sep,
	}
}

type group struct {
	items     []Node
	open      string
	close     string
	separator string
}

func (g *group) append(node ...Node) *group {
	g.items = append(g.items, node...)
	return g
}

func (g *group) render(w io.Writer) {
	if g.open != "" {
		writeString(w, g.open)
	}

	isfirst := true
	for _, node := range g.items {
		if !isfirst {
			writeString(w, g.separator)
		}
		node.render(w)
		isfirst = false
	}

	if g.close != "" {
		writeString(w, g.close)
	}
}

func (g *group) WriteTo(w io.Writer) {
	g.render(w)
}

func (g *group) WriteFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file %s: %s", path, err)
	}
	g.render(file)
	return nil
}

func (g *group) Comment(content string) *group {
	g.append(Comment(content))
	return g
}

func (g *group) CommentF(content string, args ...interface{}) *group {
	g.append(CommentF(content, args...))
	return g
}

func (g *group) Package(name string) *group {
	g.append(Package(name))
	return g
}

func (g *group) Imports() *iimport {
	i := Imports()
	g.append(i)
	return i
}

func (g *group) Line() *group {
	g.append(Line())
	return g
}

func (g *group) Var() *ivar {
	i := Var()
	g.append(i)
	return i
}

func (g *group) Const() *iconst {
	i := Const()
	g.append(i)
	return i
}

func (g *group) Function(name string) *ifunction {
	f := Function(name)
	g.append(f)
	return f
}

func (g *group) Struct(name string) *istruct {
	i := Struct(name)
	g.append(i)
	return i
}
