package gg

import "testing"

func TestType(t *testing.T) {
	t.Run("type def", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "type xstring string"

		Type("xstring", "string").render(buf)

		compareAST(t, expected, buf.String())
	})
	t.Run("type alias", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "type xstring = string"

		TypeAlias("xstring", "string").render(buf)

		compareAST(t, expected, buf.String())
	})
}
