package gg

import "fmt"

func ExampleNewGroup() {
	f := NewGroup()
	f.AddPackage("main")
	f.NewImport().AddPath("fmt")
	f.NewFunction("main").AddBody(
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
