package gg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "package test\n"

	Package("test").render(buf)

	assert.Equal(t, expected, buf.String())
}
