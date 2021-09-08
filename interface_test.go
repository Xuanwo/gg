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
	in.NewFunction("TestA").
		AddParameter("a", "int64").
		AddParameter("b", "int")
	in.NewFunction("TestB").
		AddResult("err", "error")

	in.render(buf)

	compareAST(t, expected, buf.String())
}
