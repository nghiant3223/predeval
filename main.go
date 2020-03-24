package main

import (
	"fmt"

	"github.com/ntncsebku/reloader/ast"
)

func main() {
	parser := ast.NewParser()

	result, err := parser.Parse("x + + + x")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
