package gg

import (
	"testing"
)

func TestImports(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `import (
// test
"context"
. "time"
_ "math"

test "testing"
)
`
	Import().
		AddLineComment("test").
		AddPath("context").
		AddDot("time").
		AddBlank("math").
		AddLine().
		AddAlias("testing", "test").
		render(buf)

	compareAST(t, expected, buf.String())
}
