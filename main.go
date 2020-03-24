package main

import (
	"fmt"

	"github.com/ntncsebku/reloader/ast"
	"github.com/ntncsebku/reloader/eval"
)

func main() {
	parser := ast.NewParser()
	evaluator := eval.NewEvaluator()

	tree, err := parser.Parse("1 and 2.0")
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := evaluator.Evaluate(tree, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
