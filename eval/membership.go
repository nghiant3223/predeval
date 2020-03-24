package eval

import (
	"github.com/ntncsebku/reloader/ast/expression"
	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/op"
)

func (v *visitor) evaluateMembershipOperator(leftResult interface{}, rightOperand []expression.Expression, operator op.Operator) interface{} {
	switch operator {
	case op.In:
		// Return TRUE if left result equals to one of the right item result
		for _, operandItem := range rightOperand {
			itemResult := v.Visit(operandItem)
			if areEqual, _ := v.evaluateBinaryOperator(leftResult, itemResult, op.Eq).(bool); areEqual {
				return true
			}
		}
		return false
	default:
		panic(errors.InvalidOperation)
	}
}
