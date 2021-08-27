package codegen

import "fmt"

type Const struct {
	name  string
	ty    string
	value string
}

func NewConst(name, ty, value string) *Const {
	return &Const{
		name:  name,
		ty:    ty,
		value: value,
	}
}

func (c *Const) String() string {
	return fmt.Sprintf("%s %s = %s", c.name, c.ty, c.value)
}

type ConstGroup struct {
	cs []*Const
}

func NewConstGroup() *ConstGroup {
	return &ConstGroup{}
}

func (c *ConstGroup) Add(name, ty, value string) *ConstGroup {
	c.cs = append(c.cs, NewConst(name, ty, value))
	return c
}

func (c *ConstGroup) String() string {
	switch len(c.cs) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("const %s", c.cs[0])
	default:
		buf := pool.Get()
		defer buf.Free()

		buf.AppendString("const (\n")
		for _, v := range c.cs {
			buf.AppendString(v.String())
			buf.AppendString("\n")
		}
		buf.AppendString(")\n")
		return buf.String()
	}
}
