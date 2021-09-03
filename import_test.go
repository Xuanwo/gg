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
	Imports().
		Path("context").
		Dot("time").
		Blank("math").
		Alias("testing", "test").
		render(buf)

	compareAST(t, expected, buf.String())
}
