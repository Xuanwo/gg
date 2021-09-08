package gg

import "io"

type iif struct {
	judge Node
	body  *group
}

func If(judge Node) *iif {
	return &iif{
		judge: judge,
		body:  newGroup("{\n", "\n}", "\n"),
	}
}
func (i *iif) render(w io.Writer) {
	writeString(w, "if ")
	i.judge.render(w)
	i.body.render(w)
}

func (i *iif) AddBody(node ...interface{}) *iif {
	i.body.append(node...)
	return i
}

type ifor struct {
	judge Node
	body  *group
}

func (i *ifor) render(w io.Writer) {
	writeString(w, "for ")
	i.judge.render(w)
	i.body.render(w)
}

func For(judge Node) *ifor {
	return &ifor{
		judge: judge,
		body:  newGroup("{\n", "\n}", "\n"),
	}
}

func (i *ifor) AddBody(node ...interface{}) *ifor {
	i.body.append(node...)
	return i
}

type icase struct {
	judge Node // judge == nil means it's a default case.
	body  *group
}

func (i *icase) render(w io.Writer) {
	if i.judge == nil {
		writeString(w, "default:")
	} else {
		writeString(w, "case ")
		i.judge.render(w)
		writeString(w, ":")
	}
	i.body.render(w)
}

func (i *icase) AddBody(node ...interface{}) *icase {
	i.body.append(node...)
	return i
}

type iswitch struct {
	judge       Node
	cases       []*icase
	defaultCase *icase
}

func (i *iswitch) render(w io.Writer) {
	writeString(w, "switch ")
	i.judge.render(w)
	writeString(w, "{\n")
	for _, c := range i.cases {
		c.render(w)
		writeString(w, "\n")
	}
	if i.defaultCase != nil {
		i.defaultCase.render(w)
		writeString(w, "\n")
	}
	writeString(w, "}")
}

func Switch(judge Node) *iswitch {
	return &iswitch{
		judge: judge,
	}
}
func (i *iswitch) NewCase(judge Node) *icase {
	ic := &icase{
		judge: judge,
		body:  newGroup("\n", "", "\n"),
	}
	i.cases = append(i.cases, ic)
	return ic
}

func (i *iswitch) NewDefault() *icase {
	ic := &icase{
		body: newGroup("\n", "", "\n"),
	}
	i.cases = append(i.cases, ic)
	return ic
}

func Continue() Node {
	return String("continue")
}
