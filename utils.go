package gg

import (
	"fmt"
	"io"
)

func writeString(w io.Writer, s ...string) {
	for _, v := range s {
		_, err := w.Write([]byte(v))
		if err != nil {
			panic(fmt.Errorf("write string: %v", err))
		}
	}
}

func writeStringF(w io.Writer, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	_, err := w.Write([]byte(s))
	if err != nil {
		panic(fmt.Errorf("write string: %v", err))
	}
}
