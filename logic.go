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
func (i *iif) Body(node ...Node) *iif {
	i.body.append(node...)
	return i
}

type ifor struct {
}

func (v *ifor) render(w io.Writer) {
}
func For(judge Node) *ifor {
	return nil
}
func (i *ifor) Body(node ...Node) *ifor {
	return nil
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

func (i *icase) Body(node ...Node) *icase {
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
	writeString(w, "{")
	for _, c := range i.cases {
		c.render(w)
	}
	if i.defaultCase != nil {
		i.defaultCase.render(w)
	}
	writeString(w, "}")
}

func Switch(judge Node) *iswitch {
	return &iswitch{
		judge: judge,
	}
}
func (i *iswitch) Case(judge Node) *icase {
	ic := &icase{
		judge: judge,
		body:  newGroup("\n", "", "\n"),
	}
	i.cases = append(i.cases, ic)
	return ic
}

func (i *iswitch) Default() *icase {
	ic := &icase{
		body: newGroup("\n", "", "\n"),
	}
	i.cases = append(i.cases, ic)
	return ic
}

func Continue() Node {
	return String("continue")
}
