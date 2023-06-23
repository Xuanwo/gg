package gg

import "testing"

func TestFunction(t *testing.T) {
	t.Run("no receiver", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func Test(a int, b string) (d uint)`

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

	t.Run("no name result - no receiver - single result", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func Test(a int) (int)`

		Function("Test").
			AddParameter("a", "int").
			AddResult("", "int").
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("no name result - no receiver - multi result", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func Test(a int) (int, string, error)`

		Function("Test").
			AddParameter("a", "int").
			AddResult("", "int").
			AddResult("", "string").
			AddResult("", "error").
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("no name result - has receiver - single result", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func (r *Q) Test(a int) (int)`

		Function("Test").
			WithReceiver("r", "*Q").
			AddParameter("a", "int").
			AddResult("", "int").
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("no name result - has receiver - multi result", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `func (r *Q) Test(a int) (int, string, error)`

		Function("Test").
			WithReceiver("r", "*Q").
			AddParameter("a", "int").
			AddResult("", "int").
			AddResult("", "string").
			AddResult("", "error").
			render(buf)

		compareAST(t, expected, buf.String())
	})
}
