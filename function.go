package gg

import "io"

type ifunction struct {
	comments   *group
	name       string
	receiver   Node
	parameters *group
	results    *group
	body       *group
}

// Function represent both method and function in Go.
//
// If receiver is nil, we will generate like a pure function.
// Or, we will generate a method.
func Function(name string) *ifunction {
	i := &ifunction{
		name:       name,
		comments:   newGroup("", "", "\n"),
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
	if i.comments.length() != 0 {
		i.comments.render(w)
		// We always need to insert a new line for function comments
		writeString(w, "\n")
	}

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

	// Render body
	i.body.render(w)
}

// LineComment will insert a new line comment.
func (i *ifunction) LineComment(content string, args ...interface{}) *ifunction {
	i.comments.append(LineComment(content, args...))
	return i
}

// NamedLineComment will insert a new line comment starts with function name.
func (i *ifunction) NamedLineComment(content string, args ...interface{}) *ifunction {
	content = i.name + " " + content
	i.comments.append(LineComment(content, args...))
	return i
}

func (i *ifunction) Receiver(name, typ string) *ifunction {
	i.receiver = &ifield{
		name:      name,
		value:     typ,
		separator: " ",
	}
	return i
}

func (i *ifunction) Parameter(name, typ string) *ifunction {
	i.parameters.append(&ifield{
		name:      name,
		value:     typ,
		separator: " ",
	})
	return i
}

func (i *ifunction) Result(name, typ string) *ifunction {
	i.results.append(&ifield{
		name:      name,
		value:     typ,
		separator: " ",
	})
	return i
}

func (i *ifunction) Body(node ...Node) *ifunction {
	i.body.append(node...)
	return i
}
