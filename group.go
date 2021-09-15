package gg

import (
	"fmt"
	"io"
	"os"
)

func NewGroup() *group {
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

	// NewIf this result is true, we will omit the wrap like `()`, `{}`.
	omitWrapIf func() bool
}

func (g *group) length() int {
	return len(g.items)
}

func (g *group) shouldOmitWrap() bool {
	if g.omitWrapIf == nil {
		return false
	}
	return g.omitWrapIf()
}

func (g *group) append(node ...interface{}) *group {
	if len(node) == 0 {
		return g
	}
	g.items = append(g.items, parseNodes(node)...)
	return g
}

func (g *group) render(w io.Writer) {
	if g.open != "" && !g.shouldOmitWrap() {
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

	if g.close != "" && !g.shouldOmitWrap() {
		writeString(w, g.close)
	}
}

func (g *group) Write(w io.Writer) {
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

func (g *group) AppendFile(path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("create file %s: %s", path, err)
	}
	g.render(file)
	return nil
}

func (g *group) String() string {
	buf := pool.Get()
	defer buf.Free()

	g.render(buf)
	return buf.String()
}

func (g *group) AddLineComment(content string, args ...interface{}) *group {
	g.append(LineComment(content, args...))
	return g
}

func (g *group) AddPackage(name string) *group {
	g.append(Package(name))
	return g
}

func (g *group) AddLine() *group {
	g.append(Line())
	return g
}

func (g *group) AddString(content string, args ...interface{}) *group {
	g.append(S(content, args...))
	return g
}

func (g *group) AddType(name string, typ interface{}) *group {
	g.append(Type(name, typ))
	return g
}

func (g *group) AddTypeAlias(name string, typ interface{}) *group {
	g.append(TypeAlias(name, typ))
	return g
}

func (g *group) NewImport() *iimport {
	i := Import()
	g.append(i)
	return i
}

func (g *group) NewIf(judge Node) *iif {
	i := If(judge)
	g.append(i)
	return i
}

func (g *group) NewFor(judge Node) *ifor {
	i := For(judge)
	g.append(i)
	return i
}

func (g *group) NewSwitch(judge Node) *iswitch {
	i := Switch(judge)
	g.append(i)
	return i
}

func (g *group) NewVar() *ivar {
	i := Var()
	g.append(i)
	return i
}

func (g *group) NewConst() *iconst {
	i := Const()
	g.append(i)
	return i
}

func (g *group) NewFunction(name string) *ifunction {
	f := Function(name)
	g.append(f)
	return f
}

func (g *group) NewStruct(name string) *istruct {
	i := Struct(name)
	g.append(i)
	return i
}

func (g *group) NewInterface(name string) *iinterface {
	i := Interface(name)
	g.append(i)
	return i
}
