package gg

import "fmt"

func ExampleGroup() {
	f := Group()
	f.Package("main")
	f.Imports().Path("fmt")
	f.Function("main").Body(
		String(`fmt.Println("%s")`, "Hello, World!"),
	)
	fmt.Println(f.String())
	// Output:
	// package main
	//
	// import "fmt"
	// func main(){
	// fmt.Println("Hello, World!")}
}
