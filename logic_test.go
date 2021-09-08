package gg

import "testing"

func TestIf(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `
if ok {
	println("Hello, World!")
}
`

	If(String("ok")).AddBody(String(`println("Hello, World!")`)).
		render(buf)

	compareAST(t, expected, buf.String())
}

func TestSwitch(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `
switch x {
case 1:
	print("1")
case 2:
	print("2")
default:
	print("default")
}
`
	is := Switch(String("x"))
	is.NewCase(String("1")).AddBody(String(`print("1")`))
	is.NewCase(String("2")).AddBody(String(`print("2")`))
	is.NewDefault().AddBody(String(`print("default")`))
	is.render(buf)

	compareAST(t, expected, buf.String())
}
