package expression

import (
	"fmt"

	"github.com/ntncsebku/reloader/eval"
	"github.com/ntncsebku/reloader/op"
)

type BinaryExpression struct {
	Operator     op.Operator
	RightOperand Expression
	LeftOperand  Expression
}

func NewBinaryExpression(operator string, rightOperand Expression, leftOperand Expression) *BinaryExpression {
	return &BinaryExpression{
		Operator:     op.GetOperatorByString(operator),
		RightOperand: rightOperand,
		LeftOperand:  leftOperand,
	}
}

func (e *BinaryExpression) String() string {
	return fmt.Sprintf("Binary(%s,%s,%s)", e.Operator, e.LeftOperand, e.RightOperand)
}

func (e *BinaryExpression) Accept(v eval.Visitor) {

}
