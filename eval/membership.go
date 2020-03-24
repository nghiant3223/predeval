package eval

import (
	"github.com/ntncsebku/reloader/ast/expression"
	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/op"
)

func (v *visitor) evaluateMembershipOperator(leftResult interface{}, rightOperand []expression.Expression, operator op.Operator) (interface{}, error) {
	switch operator {
	case op.In:
		// Return TRUE if left result equals to one of the right item result
		for _, operandItem := range rightOperand {
			itemResult, err := v.Visit(operandItem)
			if err != nil {
				return nil, err
			}
			isEqual, err := v.evaluateBinaryOperator(leftResult, itemResult, op.Eq)
			// If leftResult is not comparable with this item, try next item
			if err != nil && err != errors.InvalidOperation {
				return nil, err
			}
			if isEqualInBool, _ := isEqual.(bool); isEqualInBool {
				return true, nil
			}
		}
		return false, nil
	default:
		return nil, errors.InvalidOperation
	}
}
