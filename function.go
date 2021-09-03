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

func Function(name string) *ifunction {
	return &ifunction{
		name:       name,
		comments:   newGroup("", "", "\n"),
		parameters: newGroup("(", ")", ","),
		results:    newGroup("(", ")", ","),
		body:       newGroup("{\n", "}", "\n"),
	}
}

func (f *ifunction) render(w io.Writer) {
	f.comments.render(w)

	writeString(w, "func ")

	// Render receiver
	if f.receiver != nil {
		writeString(w, "(")
		f.receiver.render(w)
		writeString(w, ")")
	}

	// Render function name
	writeString(w, f.name)

	// Render parameters
	f.parameters.render(w)

	// Render results
	// FIXME: we always render `()` here, maybe we can remove it if only one result here.
	f.results.render(w)

	// Render body
	f.body.render(w)
}

func (f *ifunction) Comment(content string) *ifunction {
	f.comments.append(Comment(content))
	return f
}

func (f *ifunction) CommentF(content string, args ...interface{}) *ifunction {
	f.comments.append(CommentF(content, args...))
	return f
}

func (f *ifunction) Receiver(name, typ string) *ifunction {
	f.receiver = &ifield{
		name:      name,
		value:     typ,
		separator: " ",
	}
	return f
}

func (f *ifunction) Parameter(name, typ string) *ifunction {
	f.parameters.append(&ifield{
		name:      name,
		value:     typ,
		separator: " ",
	})
	return f
}

func (f *ifunction) Result(name, typ string) *ifunction {
	f.results.append(&ifield{
		name:      name,
		value:     typ,
		separator: " ",
	})
	return f
}

func (f *ifunction) Body(node ...Node) *ifunction {
	f.body.append(node...)
	return f
}
