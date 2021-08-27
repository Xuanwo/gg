package codegen

import (
	"fmt"
	"text/template"
)

type Function struct {
	comment    string
	receiver   *Receiver
	name       string
	parameters []*Parameter
	results    []*Result
	bodyTmpl   string
	bodyData   interface{}
}

func NewFunction(name string) *Function {
	return &Function{
		receiver: nil,
		name:     name,
	}
}

func NewMethod(r *Receiver, name string) *Function {
	return &Function{
		receiver: r,
		name:     name,
	}
}

func (f *Function) WithComment(comment string) *Function {
	f.comment = comment
	return f
}

func (f *Function) WithParameters(ps ...*Parameter) *Function {
	f.parameters = ps
	return f
}

func (f *Function) WithResults(rs ...*Result) *Function {
	f.results = rs
	return f
}

func (f *Function) WithBody(tmpl string, data interface{}) *Function {
	f.bodyTmpl = tmpl
	f.bodyData = data
	return f
}

// String will generate the function into.
//
// func (r *Receiver) FunctionName(p Parameter) (r Result) { body }
func (f *Function) String() string {
	buf := pool.Get()
	defer buf.Free()

	if len(f.comment) > 0 {
		buf.AppendString(formatComment(f.comment))
	}

	buf.AppendString("func ")
	// Generate receiver is we have one.
	if f.receiver != nil {
		buf.AppendString("(")
		buf.AppendString(f.receiver.String())
		buf.AppendString(")")
	}
	buf.AppendString(f.name)

	// Generate method parameters
	buf.AppendString("(")
	// FIXME: we need to support rewrite `x int, y int` into `x, y int`.
	isFirst := true
	for _, v := range f.parameters {
		if !isFirst {
			buf.AppendString(",")
		}
		buf.AppendString(v.String())
		isFirst = false
	}
	buf.AppendString(")")

	// Generate method results
	buf.AppendString("(")
	// FIXME: we need to support rewrite `x int, y int` into `x, y int`.
	isFirst = true
	for _, v := range f.results {
		if !isFirst {
			buf.AppendString(",")
		}
		buf.AppendString(v.String())
		isFirst = false
	}
	buf.AppendString(")")
	// Generate method body.
	buf.AppendString(" {\n")
	tmpl, err := template.New(f.name).Parse(f.bodyTmpl)
	if err != nil {
		panic(fmt.Errorf("parse template: %w", err))
	}
	err = tmpl.Execute(buf, f.bodyData)
	if err != nil {
		panic(fmt.Errorf("execute template: %w", err))
	}
	buf.AppendString("}\n")

	return buf.String()
}

type Receiver struct {
	name      string
	ty        string
	isPointer bool
}

func NewReceiver(name, ty string, isPointer bool) *Receiver {
	return &Receiver{
		name:      name,
		ty:        ty,
		isPointer: isPointer,
	}
}

func (r *Receiver) String() string {
	buf := pool.Get()
	defer buf.Free()

	buf.AppendString(r.name)
	buf.AppendString(" ")
	if r.isPointer {
		buf.AppendString(" *")
	}
	buf.AppendString(r.ty)
	return buf.String()
}

type Parameter struct {
	name      string
	ty        string
	isPointer bool
}

func NewParameter(name, ty string, isPointer bool) *Parameter {
	return &Parameter{
		name:      name,
		ty:        ty,
		isPointer: isPointer,
	}
}

func (r *Parameter) String() string {
	buf := pool.Get()
	defer buf.Free()

	buf.AppendString(r.name)
	buf.AppendString(" ")
	if r.isPointer {
		buf.AppendString("*")
	}
	buf.AppendString(r.ty)
	return buf.String()
}

type Result struct {
	name      string
	ty        string
	isPointer bool
}

func NewResult(name, ty string, isPointer bool) *Result {
	return &Result{
		name:      name,
		ty:        ty,
		isPointer: isPointer,
	}
}

func (r *Result) String() string {
	buf := pool.Get()
	defer buf.Free()

	buf.AppendString(r.name)
	buf.AppendString(" ")
	if r.isPointer {
		buf.AppendString("*")
	}
	buf.AppendString(r.ty)
	return buf.String()
}
