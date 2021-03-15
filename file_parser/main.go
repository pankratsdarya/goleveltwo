package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

const (
	srcFileName  = "src.go"
	functionName = "lockUnlock"
)

func main() {
	funcCount, err := getCount(srcFileName, functionName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s count: %d", functionName, funcCount)
}

func getCount(srcFileName, functionName string) (int, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, srcFileName, nil, 0)
	if err != nil {
		return 0, err
	}

	var fCount int
	ast.Inspect(astFile, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.Ident:
			s = x.Name
		}
		if s == functionName {
			fCount++
		}
		return true
	})

	return fCount, nil
}
