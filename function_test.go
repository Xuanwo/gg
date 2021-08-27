package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunction_String(t *testing.T) {
	m := &Function{
		comment:  "Test function comment",
		receiver: NewReceiver("t", "TestFunctionReceiver", false),
		name:     "TestFunctionName",
		parameters: []*Parameter{
			NewParameter("p1", "TestFunctionParameter", true),
			NewParameter("p2", "int64", false),
		},
		results: []*Result{
			NewResult("r1", "TestFunctionResult", true),
			NewResult("r2", "int64", false),
		},
		bodyTmpl: `println("{{.}}")`,
		bodyData: "Test Body",
	}

	expect := `// Test function comment
func (t TestFunctionReceiver)TestFunctionName(p1 *TestFunctionParameter,p2 int64)(r1 *TestFunctionResult,r2 int64) {
println("Test Body")}
`

	assert.Equal(t, expect, m.String())
}
