package gg

import (
	"fmt"
	"io"
)

type icall struct {
	owner Node
	name  string
	items *group
	calls *group
}

// Call is used to generate a function call.
func Call(name string) *icall {
	ic := &icall{
		name:  name,
		items: newGroup("(", ")", ","),
		calls: newGroup("", "", "."),
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
	if i.calls.length() != 0 {
		writeString(w, ".")
		i.calls.render(w)
	}
}

func (i *icall) WithOwner(name string) *icall {
	if i.owner != nil {
		panic(fmt.Errorf("icall already have owner %v", i.owner))
	}
	i.owner = String(name)
	return i
}

func (i *icall) AddParameter(value ...interface{}) *icall {
	i.items.append(value...)
	return i
}

func (i *icall) AddCall(name string, params ...interface{}) *icall {
	i.calls.append(Call(name).AddParameter(params...))
	return i
}
