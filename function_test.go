package gg

import "testing"

func TestFunction(t *testing.T) {
	t.Run("no receiver", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func Test(a int, b string) (d uint) {}`

		Function("Test").
			Parameter("a", "int").
			Parameter("b", "string").
			Result("d", "uint").
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("has receiver", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func (r *Q) Test() (a int, b int64, d string) {
return "Hello, World!"
}`
		Function("Test").
			Receiver("r", "*Q").
			Result("a", "int").
			Result("b", "int64").
			Result("d", "string").
			Body(
				String(`return "Hello, World!"`),
			).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("node input", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func (r *Q) Test() (a int, b int64, d string) {
return "Hello, World!"
}`
		Function("Test").
			Receiver("r", "*Q").
			Result("a", String("int")).
			Result("b", "int64").
			Result("d", "string").
			Body(
				String(`return "Hello, World!"`),
			).
			render(buf)

		compareAST(t, expected, buf.String())
	})
}
