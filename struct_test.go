package gg

import "testing"

func TestStruct(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "type Test struct{}"

		Struct("Test").render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("fields", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `type Test struct{
A int64
b string
}`

		Struct("Test").
			AddField("A", "int64").
			AddField("b", "string").
			render(buf)

		compareAST(t, expected, buf.String())
	})
}
