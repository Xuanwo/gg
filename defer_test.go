package gg

import "testing"

func TestDefer(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "defer hello()"

	Defer("hello()").render(buf)

	compareAST(t, expected, buf.String())
}
