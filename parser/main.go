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
	//for debug
	//ast.Print(fset, node)

	//listFunc(node, listFuncOption{printInOutArgument: false})
	fmt.Println("function")
	listFunc(node, listFuncOption{printInOutArgument: true})

	fmt.Println("struct")
	listStruct(node)

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

func listStruct(node *ast.File) {
	ast.Inspect(node, func(n ast.Node) bool {
		if value, ok := n.(*ast.GenDecl); ok && value.Tok == token.TYPE {
			for _, spec := range value.Specs {
				tspec, _ := spec.(*ast.TypeSpec)
				fmt.Printf("Name:%v\n", tspec.Name.Name)
				stype, _ := tspec.Type.(*ast.StructType)
				for _, field := range stype.Fields.List {
					switch fieldType := field.Type.(type) {
					case *ast.Ident:
						fmt.Printf("- filed:%v(%v)\n", field.Names[0].Name, fieldType.Name)
					default:
						panic("not implemnted for this type")
					}

				}
			}
		}
		return true
	})

}
