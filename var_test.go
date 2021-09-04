package gg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVar(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "var Version=2"

		Var().
			Field("Version", Lit(2)).
			render(buf)

		assert.Equal(t, expected, buf.String())
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
}
