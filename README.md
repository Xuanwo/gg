[![Build Status](https://github.com/Xuanwo/gg/workflows/Unit%20Test/badge.svg?branch=master)](https://github.com/Xuanwo/gg/actions?query=workflow%3A%22Unit+Test%22)
[![Go dev](https://pkg.go.dev/badge/github.com/Xuanwo/gg)](https://pkg.go.dev/github.com/Xuanwo/gg)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/Xuanwo/gg/blob/master/LICENSE)
[![matrix](https://img.shields.io/matrix/xuanwo@gg:matrix.org.svg?logo=matrix)](https://matrix.to/#/#xuanwo@gg:matrix.org)

# gg

`gg` is a General Golang Code Generator: A Good Game to play with Golang.

```go
package main

import (
	"fmt"

	. "github.com/Xuanwo/gg"
)

func main() {
	f := Group()
	f.Package("main")
	f.Imports().Path("fmt")
	f.Function("main").Body(
		String(`fmt.Println("%s")`, "Hello, World!"),
	)
	fmt.Println(f.String())
}
```

Output (after `go fmt`)

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Design

`gg` is a general golang code generator that designed for resolving problems exists in the following tools:

- [text/template](https://pkg.go.dev/text/template): Additional syntax, Steep learning curve, Complex logic is difficult to maintain
- [dave/jennifer](https://github.com/dave/jennifer): Overly abstract APIs, user need to take care about `()`, `,` everywhere.
- [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder): Parse data from struct tags/comments, not a general code generator.

In short, `gg` will provide near-native golang syntax and helpful API so play a good game with Golang. With `gg`, we can generate golang code more easily and understandable.

## Usage

### Package Name

```go
f := Group()
f.Package("main")
// package main
```

### Imports

```go
f := Group()
f.Imports().
    Path("context").
	Dot("math").
	Blank("time").
	Alias("x", "testing")
// import (
//      "context"
//      . "math"
//      _ "time"
//      x "testing"
// )
```

### Function

```go
f := Group()
f.Function("hello").
    Receiver("v", "*World").
    Parameter("content", "string").
    Parameter("times", "int").
	Result("v", "string").
	Body(gg.String(`return fmt.Sprintf("say %s in %d times", content, times)`))
// func (v *World) hello(content string, times int) (v string) {
//  return fmt.Sprintf("say %s in %d times", content, times)
//}
```

### Struct

```go
f := Group()
f.Struct("World").
	Field("x", "int64").
	Field("y", "string")
// type World struct {
//    x int64
//    y string
//}
```

## Acknowledgement

- `gg` is inspired by [dave/jennifer](https://github.com/dave/jennifer), I borrowed most ideas and some code from it. Nice work!
