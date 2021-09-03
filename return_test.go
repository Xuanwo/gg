package gg

import "testing"

func TestReturn(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `return a, b,123,Test{Abc:123}`

	ir := Return().
		Id("a").
		Id("b").
		Lit(123)
	ir.Value("Test").
		Field("Abc", "123")
	ir.render(buf)

	compareAST(t, expected, buf.String())
}
