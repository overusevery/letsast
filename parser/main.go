package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "../example_program/main.go", nil, parser.AllErrors)
	if err != nil {
		fmt.Println("parse error:", err)
		os.Exit(1)
	}

	listFuncName(node)
	//ast.Print(fset, node)
}

func listFuncName(node *ast.File) {
	ast.Inspect(node, func(n ast.Node) bool {
		value, ok := n.(*ast.FuncDecl)
		if ok {
			fmt.Printf("Name:%v\n", value.Name.Name)
		}
		return true
	})

}
