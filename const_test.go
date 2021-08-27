package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstGroup_String(t *testing.T) {
	c := NewConstGroup()

	c.Add("test", "int64", "123")
	assert.Equal(t, "const test int64 = 123", c.String())

	c.Add("test1", "int32", "456")
	assert.Equal(t, `const (
test int64 = 123
test1 int32 = 456
)
`, c.String())
}
