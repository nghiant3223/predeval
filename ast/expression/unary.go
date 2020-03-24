package expression

import (
	"fmt"

	"github.com/ntncsebku/reloader/op"
)

type UnaryExpression struct {
	Operator op.Operator
	Operand  Expression
}

func NewUnaryExpression(operator string, operand Expression) *UnaryExpression {
	return &UnaryExpression{
		Operator: op.GetOperatorByString(operator),
		Operand:  operand,
	}
}

func (e *UnaryExpression) String() string {
	return fmt.Sprintf("Unary(%s,%s)", e.Operator, e.Operand)
}
