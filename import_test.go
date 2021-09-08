package gg

import (
	"testing"
)

func TestImports(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `import (
"context"
. "time"
_ "math"
test "testing"
)
`
	Import().
		AddPath("context").
		AddDot("time").
		AddBlank("math").
		AddAlias("testing", "test").
		render(buf)

	compareAST(t, expected, buf.String())
}
