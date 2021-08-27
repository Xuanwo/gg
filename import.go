package codegen

import (
	"fmt"
)

const (
	ImportDot   = "."
	ImportBlank = "_"
)

type Import struct {
	path  string
	alias string
}

func newImport(path, alias string) *Import {
	return &Import{
		path:  path,
		alias: alias,
	}
}

func NewImport(path string) *Import {
	return newImport(path, "")
}

func NewDotImport(path string) *Import {
	return newImport(path, ImportDot)
}

func NewBlankImport(path string) *Import {
	return newImport(path, ImportBlank)
}

func (i *Import) String() string {
	return fmt.Sprintf(`%s "%s"`, i.alias, i.path)
}

type ImportGroup struct {
	ims []*Import
}

func NewImportGroup() *ImportGroup {
	return &ImportGroup{}
}

func (is *ImportGroup) Add(ims ...*Import) *ImportGroup {
	is.ims = append(is.ims, ims...)
	return is
}

// TODO: we need to add test for this.
func (is *ImportGroup) String() string {
	switch len(is.ims) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("import %s", is.ims[0])
	default:
		buf := pool.Get()
		defer buf.Free()

		buf.AppendString("import (\n")
		for _, v := range is.ims {
			buf.AppendString(v.String())
			buf.AppendString("\n")
		}
		buf.AppendString(")\n")
		return buf.String()
	}
}
