package gg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"testing"
	"text/template"
)

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

func TestParseAST(t *testing.T) {
	content := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.AllErrors)
	if err != nil {
		t.Error(err)
	}
	ast.Print(fset, f)
}

func TestViaGolangAST(t *testing.T) {
	fset := token.NewFileSet()
	f := &ast.File{
		Name:  ast.NewIdent("main"),
		Scope: ast.NewScope(nil),
	}

	f.Decls = append(f.Decls, &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"fmt"`,
				},
			},
		},
	})

	f.Decls = append(f.Decls, &ast.FuncDecl{
		Name: ast.NewIdent("main"),
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("fmt"),
						Sel: ast.NewIdent("Println"),
					},
					Args: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: `"Hello, World!"`,
						},
					},
				}},
			},
		},
	})

	err := format.Node(os.Stdout, fset, f)
	if err != nil {
		log.Fatalf("ast is incorrect")
	}
}

func TestViaString(t *testing.T) {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, "package %s\n\n", "main")
	fmt.Fprintf(b, "import %s\n\n", `"fmt"`)
	fmt.Fprintf(b, "func main() {\n")
	fmt.Fprintf(b, "\tfmt.Println(%s)\n", `"Hello, World!"`)
	fmt.Fprint(b, "}\n")

	fmt.Println(b.String())
}

func TestViaGolangTemplate(t *testing.T) {
	b := &bytes.Buffer{}

	data := struct {
		Package string
		Import  []string
		Content string
	}{
		Package: "main",
		Import:  []string{"fmt"},
		Content: "Hello, World!",
	}

	tmpl := `package {{ .Package }}

import (
	{{ range $_, $v := .Import -}}
		"{{ $v }}"
	{{ end -}}
)

func main() {
	fmt.Println("{{ .Content }}")
}`

	err := template.Must(template.New("test").Parse(tmpl)).Execute(b, data)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(b.String())
}
