package gg

import "io"

type ifunction struct {
	name       string
	receiver   Node
	parameters *group
	results    *group
	body       *group
	call       *icall
}

// Function represent both method and function in Go.
//
// NOTES
//
// If `WithReceiver`, we will generate a method:
//    func (t test) Test()
//
// If `WithCall`, we will generate a function call:
//    func Test(){}()
//
// If `AddBody`, we will generate like a function definition without body:
//    func Test() {
//        println("Hello, World!")
//    }
func Function(name string) *ifunction {
	i := &ifunction{
		name:       name,
		parameters: newGroup("(", ")", ","),
		results:    newGroup("(", ")", ","),
		body:       newGroup("{\n", "}", "\n"),
	}
	// We should omit the `()` if result is empty
	i.results.omitWrapIf = func() bool {
		l := i.results.length()
		if l == 0 {
			// There is no result fields, we can omit `()` safely.
			return true
		}
		// NOTE: We also need to omit `()` while there is only one field,
		//  and the field name is empty, like `test() (int64) => test() int64`.
		//  But it's hard to implement in render side, so we let `go fmt` to do the job.
		return false
	}
	return i
}

func (i *ifunction) render(w io.Writer) {
	writeString(w, "func ")

	// Render receiver
	if i.receiver != nil {
		writeString(w, "(")
		i.receiver.render(w)
		writeString(w, ")")
	}

	// Render function name
	writeString(w, i.name)

	// Render parameters
	i.parameters.render(w)

	// Render results
	i.results.render(w)

	// Only render body while there is a body or a call.
	//
	// This will add extra burden for functions that have empty body.
	// But it's a rare case, and we can always add an empty line in body to workaround.
	if i.body.length() > 0 || i.call != nil {
		i.body.render(w)
	}

	// Only render function call while there is a call.
	if i.call != nil {
		i.call.render(w)
	}
}

func (i *ifunction) WithReceiver(name, typ interface{}) *ifunction {
	i.receiver = field(name, typ, " ")
	return i
}

func (i *ifunction) WithCall(params ...interface{}) *ifunction {
	i.call = Call("").AddParameter(params...)
	return i
}

func (i *ifunction) AddParameter(name, typ interface{}) *ifunction {
	i.parameters.append(field(name, typ, " "))
	return i
}

func (i *ifunction) AddResult(name, typ interface{}) *ifunction {
	i.results.append(field(name, typ, " "))
	return i
}

func (i *ifunction) AddBody(node ...interface{}) *ifunction {
	i.body.append(node...)
	return i
}
