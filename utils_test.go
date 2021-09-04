package gg

import (
	"strings"
	"testing"
)

// cleanAST will remove all space and newline from code.
// We know it will break the AST semantic, but golang doesn't support parse
// partial code source, we have to do like this.
// Maybe we can find a better way to compare the AST in go.
func cleanAST(a string) string {
	a = strings.ReplaceAll(a, " ", "")
	a = strings.ReplaceAll(a, "\n", "")
	a = strings.ReplaceAll(a, "\t", "")
	return a
}

func compareAST(t *testing.T, a, b string) {
	na, nb := cleanAST(a), cleanAST(b)
	if na == nb {
		return
	}
	t.Error("AST is not the same.")
	t.Errorf("left:\n%s\ncleaned:\n%s", a, na)
	t.Errorf("right:\n%s\ncleaned:\n%s", b, nb)
}
