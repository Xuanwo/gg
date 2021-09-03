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

	If(String("ok")).Body(String(`println("Hello, World!")`)).
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
	is.Case(String("1")).Body(String(`print("1")`))
	is.Case(String("2")).Body(String(`print("2")`))
	is.Default().Body(String(`print("default")`))
	is.render(buf)

	compareAST(t, expected, buf.String())
}
