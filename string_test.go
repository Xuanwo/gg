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

	expected := "// Hello, World!"

	Comment("Hello, World!").render(buf)

	assert.Equal(t, expected, buf.String())
}

func TestCommentF(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "// Hello, World!"

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

func TestFormatComment(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{
			"short line",
			"Value comment",
			"// Value comment",
		},
		{
			"long single line",
			"These is a long line that we need to do line break at 140. However, this long line is not long enough, so we still need to pollute a lot water in it. After all these jobs, we can test this long line.",
			`// These is a long line that we need to do line break at 140. However, this long line is not long enough,
// so we still need to pollute a lot water in it. After all these jobs, we can test this long line.`,
		},
		{
			"multi lines",
			`There is line which has
its own line break.`,
			`// There is line which has
// its own line break.`,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, v.expect, formatComment(v.input))
		})
	}
}
