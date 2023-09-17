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

	//listFunc(node, listFuncOption{printInOutArgument: false})
	listFunc(node, listFuncOption{printInOutArgument: true})

	//for debug
	//ast.Print(fset, node)
}

type listFuncOption struct {
	printInOutArgument bool
}

func listFunc(node *ast.File, opt listFuncOption) {
	ast.Inspect(node, func(n ast.Node) bool {
		value, ok := n.(*ast.FuncDecl)
		if ok {
			fmt.Printf("Name:%v\n", value.Name.Name)
			if opt.printInOutArgument {
				for _, param := range value.Type.Params.List {
					switch paramType := param.Type.(type) {
					case *ast.Ident:
						for _, variable := range param.Names {
							fmt.Printf(" - input:%v(%v)\n", variable.Name, paramType)
						}
					case *ast.Ellipsis:
						//example: (int...)
						for _, variable := range param.Names {
							ident, _ := paramType.Elt.(*ast.Ident)
							fmt.Printf(" - input:%v(%v...)\n", variable.Name, ident.Name)
						}
					default:
						panic("not implemnted for this type")
					}

				}
				if value.Type.Results != nil {
					for _, output := range value.Type.Results.List {
						fmt.Printf(" - output:(%v)\n", output.Type)
					}
				}
			}
		}
		return true
	})

}
