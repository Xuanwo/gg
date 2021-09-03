package gg

import (
	"strings"

	"github.com/Xuanwo/go-bufferpool"
)

var pool = bufferpool.New(1024)

// TODO: we will support use to config this logic.
const lineLength = 80

func formatComment(comment string) string {
	buf := pool.Get()
	defer buf.Free()

	// Trim space before going further.
	comment = strings.TrimSpace(comment)

	// Split comment into lines (we will keep original line break.)
	lines := strings.Split(comment, "\n")

	for _, line := range lines {
		cur := 0

		// Start a comment line.
		buf.AppendString("//")

		// Split comment into words
		words := strings.Split(line, " ")

		for _, word := range words {
			// If current line is long enough we need to break it.
			if cur >= lineLength {
				buf.AppendString("\n//")
				cur = 0
			}

			buf.AppendString(" ")
			buf.AppendString(word)
			cur += len(word)
		}
		buf.AppendString("\n")
	}

	return buf.String()
}
