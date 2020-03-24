package eval

import (
	"reflect"

	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/op"
)

//// Evaluate for op Or And Eq Neq Le Lte Ge Gte  Add Sub Mul Div Mod
func (v *visitor) evaluateBinaryOperator(leftResult, rightResult interface{}, operator op.Operator) (interface{}, error) {
	leftResult, rightResult, err := assureBinaryType(leftResult, rightResult)
	if err != nil {
		return nil, err
	}

	if leftInt, ok := leftResult.(int64); ok {
		rightInt, _ := rightResult.(int64)
		return evaluateBinaryInt(leftInt, rightInt, operator)
	}

	if leftFloat, ok := leftResult.(float64); ok {
		rightFloat, _ := rightResult.(float64)
		return evaluateBinaryFloat(leftFloat, rightFloat, operator)
	}

	if leftString, ok := leftResult.(string); ok {
		rightString, _ := rightResult.(string)
		return evaluateBinaryString(leftString, rightString, operator)
	}

	if leftBool, ok := leftResult.(bool); ok {
		rightBool, _ := rightResult.(bool)
		return evaluateBinaryBool(leftBool, rightBool, operator)
	}

	return nil, errors.InvalidDataType
}

func assureBinaryType(left, right interface{}) (coercedLeft interface{}, coercedRight interface{}, err error) {
	// If one of the operand is float64, cast the other to float64
	leftInt, isLeftInt := left.(int64)
	rightFloat, isRightFloat := right.(float64)
	if isLeftInt && isRightFloat {
		return float64(leftInt), rightFloat, nil
	}

	// If one of the operand is float64, cast the other to float64
	leftFloat, isLeftFloat := left.(float64)
	rightInt, isRightInt := right.(int64)
	if isLeftFloat && isRightInt {
		return leftFloat, float64(rightInt), nil
	}

	// Reject for other cases when type of left operand is different from that of right operand
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, nil, errors.InvalidOperation
	}

	return left, right, nil
}

func evaluateBinaryInt(left, right int64, operator op.Operator) (interface{}, error) {
	switch operator {
	// Logical
	case op.Or:
		return left != 0 || right != 0, nil
	case op.And:
		return left != 0 && right != 0, nil
	// Comparision
	case op.Eq:
		return left == right, nil
	case op.Neq:
		return left != right, nil
	case op.Le:
		return left < right, nil
	case op.Ge:
		return left > right, nil
	case op.Lte:
		return left <= right, nil
	case op.Gte:
		return left >= right, nil
	// Arithmetic
	case op.Add:
		return left + right, nil
	case op.Sub:
		return left - right, nil
	case op.Mul:
		return left * right, nil
	case op.Div:
		if right == 0 {
			return nil, errors.DivideByZero
		}
		return left / right, nil
	case op.Mod:
		return left % right, nil
	default:
		return nil, errors.InvalidOperation
	}
}

func evaluateBinaryFloat(left, right float64, operator op.Operator) (interface{}, error) {
	switch operator {
	// Logical
	case op.Or:
		return left != 0 || right != 0, nil
	case op.And:
		return left != 0 && right != 0, nil
	// Comparision
	case op.Eq:
		return left == right, nil
	case op.Neq:
		return left != right, nil
	case op.Le:
		return left < right, nil
	case op.Ge:
		return left > right, nil
	case op.Lte:
		return left <= right, nil
	case op.Gte:
		return left >= right, nil
	// Arithmetic
	case op.Add:
		return left + right, nil
	case op.Sub:
		return left - right, nil
	case op.Mul:
		return left * right, nil
	case op.Div:
		if right == 0 {
			return nil, errors.DivideByZero
		}
		return left / right, nil
	default:
		return nil, errors.InvalidOperation
	}
}

func evaluateBinaryString(left, right string, operator op.Operator) (interface{}, error) {
	switch operator {
	// Logical
	case op.Or:
		return left != "" || right != "", nil
	case op.And:
		return left != "" && right != "", nil
	// Comparision
	case op.Eq:
		return left == right, nil
	case op.Neq:
		return left != right, nil
	case op.Le:
		return left < right, nil
	case op.Ge:
		return left > right, nil
	case op.Lte:
		return left <= right, nil
	case op.Gte:
		return left >= right, nil
	// Concatenation
	case op.Add:
		return left + right, nil
	default:
		return nil, errors.InvalidOperation
	}
}

func evaluateBinaryBool(left, right bool, operator op.Operator) (interface{}, error) {
	switch operator {
	// Logical
	case op.Or:
		return left || right, nil
	case op.And:
		return left && right, nil
	// Comparision
	case op.Eq:
		return left == right, nil
	case op.Neq:
		return left != right, nil
	default:
		return nil, errors.InvalidOperation
	}
}
