package gg

import "testing"

func TestCalls(t *testing.T) {
	t.Run("no owner", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "List()"

		Call("List").render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("witch owner", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "x.List()"

		Call("List").
			Owner("x").
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("call list", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "x.List().Next(src,dst)"

		Call("List").
			Owner("x").
			Call("Next").
			Parameter("src").
			Parameter("dst").
			render(buf)

		compareAST(t, expected, buf.String())
	})
}
