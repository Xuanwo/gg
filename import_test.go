package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport_String(t *testing.T) {
	cases := []struct {
		name   string
		path   string
		alias  string
		expect string
	}{
		{
			"normal", "math", "",
			` "math"`,
		},
		{
			"dot import", "math", ImportDot,
			`. "math"`,
		},
		{
			"black import", "math", ImportBlank,
			`_ "math"`,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			i := newImport(v.path, v.alias)

			assert.Equal(t, v.expect, i.String())
		})
	}
}
