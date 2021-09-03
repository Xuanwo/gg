package gg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "Hello, World!"

	String(expected).render(buf)

	assert.Equal(t, expected, buf.String())
}

func TestStringF(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "Hello, World!"

	StringF("Hello, %s!", "World").render(buf)

	assert.Equal(t, expected, buf.String())
}

func TestComment(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "// Hello, World!\n"

	Comment("Hello, World!").render(buf)

	assert.Equal(t, expected, buf.String())
}

func TestCommentF(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "// Hello, World!\n"

	CommentF("Hello, %s!", "World").render(buf)

	assert.Equal(t, expected, buf.String())
}

func TestLit(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "true"

	Lit(true).render(buf)

	assert.Equal(t, expected, buf.String())
}
