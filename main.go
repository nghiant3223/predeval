package main

import (
	"fmt"

	"github.com/ntncsebku/reloader/ast"
)

func main() {
	parser := ast.NewParser()

	result, err := parser.Parse("(!a+b) * (c+d)")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
