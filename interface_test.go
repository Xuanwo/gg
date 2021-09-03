package gg

import "testing"

func TestInterface(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `type Tester interface {
TestA(a int64, b int)
TestB() (err error)
}
`
	in := Interface("Tester")
	in.Function("TestA").
		Parameter("a", "int64").
		Parameter("b", "int")
	in.Function("TestB").
		Result("err", "error")

	in.render(buf)

	compareAST(t, expected, buf.String())
}
