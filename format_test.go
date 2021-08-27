package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatComment(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{
			"short line",
			"Struct comment",
			"// Struct comment\n",
		},
		{
			"long single line",
			"These is a long line that we need to do line break at 140. However, this long line is not long enough, so we still need to pollute a lot water in it. After all these jobs, we can test this long line.",
			`// These is a long line that we need to do line break at 140. However, this long line is not long enough,
// so we still need to pollute a lot water in it. After all these jobs, we can test this long line.
`,
		},
		{
			"multi lines",
			`There is line which has
its own line break.`,
			`// There is line which has
// its own line break.
`,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, v.expect, formatComment(v.input))
		})
	}
}
