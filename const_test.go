package gg

import (
	"testing"
)

func TestConst(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "const Version=2"

		Const().
			AddField("Version", Lit(2)).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("typed", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "const Version int =2"

		Const().
			AddTypedField("Version", "int", Lit(2)).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("multiple", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `const (
Version=2
Description="Hello, World!"
)
`

		Const().
			AddField("Version", Lit(2)).
			AddField("Description", Lit("Hello, World!")).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("line comment", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `const (
// Version is the version
Version=2
)
`

		Const().
			AddLineComment("Version is the version").
			AddField("Version", Lit(2)).
			render(buf)

		compareAST(t, expected, buf.String())
	})
}
