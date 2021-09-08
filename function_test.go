package gg

import "testing"

func TestFunction(t *testing.T) {
	t.Run("no receiver", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func Test(a int, b string) (d uint) {}`

		Function("Test").
			AddParameter("a", "int").
			AddParameter("b", "string").
			AddResult("d", "uint").
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
			WithReceiver("r", "*Q").
			AddResult("a", "int").
			AddResult("b", "int64").
			AddResult("d", "string").
			AddBody(
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
			WithReceiver("r", "*Q").
			AddResult("a", String("int")).
			AddResult("b", "int64").
			AddResult("d", "string").
			AddBody(
				String(`return "Hello, World!"`),
			).
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("call", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func(){}()`

		fn := Function("")
		fn.WithCall()
		fn.render(buf)

		compareAST(t, expected, buf.String())
	})
}
