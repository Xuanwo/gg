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

		expected := "x.List(src)"

		Call("List").
			WithOwner("x").
			AddParameter("src").
			render(buf)

		compareAST(t, expected, buf.String())
	})

	t.Run("call list", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "x.List().Next(src,dst)"

		Call("List").
			WithOwner("x").
			AddCall("Next", "src", "dst").
			render(buf)

		compareAST(t, expected, buf.String())
	})
}
