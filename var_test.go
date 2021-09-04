package gg

import (
	"testing"
)

func TestVar(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "var Version=2"

		Var().
			Field("Version", Lit(2)).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("typed", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "var Version int =2"

		Var().
			TypedField("Version", "int", Lit(2)).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("multiple", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `var (
Version=2
Description="Hello, World!"
)
`

		Var().
			Field("Version", Lit(2)).
			Field("Description", Lit("Hello, World!")).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("decl", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `var _ io.Reader`

		Var().
			Decl("_", "io.Reader").
			render(buf)

		compareAST(t, expected, buf.String())
	})
}
