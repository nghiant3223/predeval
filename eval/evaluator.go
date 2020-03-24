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

func (e evaluator) Evaluate(exp expression.Expression, args map[string]string) (value interface{}, err error) {
	defer func() {
		// If there is any error while evaluate predicate at runtime
		// Include type mismatch, divide by zero, ...
		if r := recover(); r != nil {
			value = nil
			err, _ = r.(error)
		}
	}()

	visitor := newEvalVisitor(args)
	value = visitor.Visit(exp)

	return value, nil
}
