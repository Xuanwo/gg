package gg

import (
	"fmt"
	"io"
)

type icall struct {
	owner Node
	name  string
	items *group
}

// Call is used to generate a function call.
func Call(name string) *icall {
	ic := &icall{
		name:  name,
		items: newGroup("(", ")", ","),
	}
	return ic
}

func (i *icall) render(w io.Writer) {
	if i.owner != nil {
		i.owner.render(w)
		writeString(w, ".")
	}
	writeString(w, i.name)
	i.items.render(w)
}

func (i *icall) Owner(name string) *icall {
	if i.owner != nil {
		panic(fmt.Errorf("icall already have owner %v", i.owner))
	}
	i.owner = String(name)
	return i
}

func (i *icall) Parameter(value interface{}) *icall {
	i.items.append(parseNode(value))
	return i
}

func (i *icall) Call(name string) *icall {
	ni := Call(name)
	// We link the current icall to the new icall's owner, so we can
	// generate the function call list.
	ni.owner = i
	return ni
}
