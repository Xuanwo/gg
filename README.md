# gg

`gg` is a General Golang Code Generator: A Good Game to play with Golang.

```go
package main

import (
	"fmt"
	
	"github.com/Xuanwo/gg"
)

func main() {
	f := Group()
	f.Package("main")
	f.Imports().Path("fmt")
	f.Function("main").Body(
		StringF(`fmt.Println("%s")`, "Hello, World!"),
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

## Acknowledgement

- `gg` is inspired by [dave/jennifer](https://github.com/dave/jennifer), I borrowed most ideas and some code from it. Nice work!
