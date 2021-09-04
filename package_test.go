package gg

import (
	"testing"
)

func TestPackage(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "package test\n"

	Package("test").render(buf)

	compareAST(t, expected, buf.String())
}
