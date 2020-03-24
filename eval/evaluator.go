package eval

import (
	"github.com/ntncsebku/reloader/ast/expression"
)

type Evaluator interface {
	Evaluate(exp expression.Expression, args map[string]string) (value interface{}, err error)
}

type evaluator struct{}

func NewEvaluator() Evaluator {
	return &evaluator{}
}

func (e evaluator) Evaluate(exp expression.Expression, args map[string]string) (interface{}, error) {
	visitor := newEvalVisitor(args)
	value, err := visitor.Visit(exp)
	if err != nil {
		return nil, err
	}

	return value, nil
}
