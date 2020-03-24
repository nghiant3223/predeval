package eval

import (
	"reflect"

	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/op"
)

//// Evaluate for op Or And Eq Neq Le Lte Ge Gte  Add Sub Mul Div Mod
func EvaluateBinaryOperator(left, right interface{}, operator op.Operator) interface{} {
	left, right = assureBinaryType(left, right)

	if leftInt, ok := left.(int64); ok {
		rightInt, _ := right.(int64)
		return evaluateBinaryInt(leftInt, rightInt, operator)
	}

	if leftFloat, ok := left.(float64); ok {
		rightFloat, _ := right.(float64)
		return evaluateBinaryFloat(leftFloat, rightFloat, operator)
	}

	if leftString, ok := left.(string); ok {
		rightString, _ := right.(string)
		return evaluateBinaryString(leftString, rightString, operator)
	}

	if leftBool, ok := left.(bool); ok {
		rightBool, _ := right.(bool)
		return evaluateBinaryBool(leftBool, rightBool, operator)
	}

	panic(errors.InvalidDataType)
}

func assureBinaryType(left, right interface{}) (coercedLeft interface{}, coercedRight interface{}) {
	// If one of the operand is float64, cast the other to float64
	leftInt, isLeftInt := left.(int64)
	rightFloat, isRightFloat := right.(float64)
	if isLeftInt && isRightFloat {
		return float64(leftInt), rightFloat
	}

	// If one of the operand is float64, cast the other to float64
	leftFloat, isLeftFloat := left.(float64)
	rightInt, isRightInt := right.(int64)
	if isLeftFloat && isRightInt {
		return leftFloat, float64(rightInt)
	}

	// Reject for other cases when type of left operand is different from that of right operand
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		panic(errors.TypeMismatch)
	}

	return left, right
}

func evaluateBinaryInt(left, right int64, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Or:
		return left != 0 || right != 0
	case op.And:
		return left != 0 && right != 0
	// Comparision
	case op.Eq:
		return left == right
	case op.Neq:
		return left != right
	case op.Le:
		return left < right
	case op.Ge:
		return left > right
	case op.Lte:
		return left <= right
	case op.Gte:
		return left >= right
	// Arithmetic
	case op.Add:
		return left + right
	case op.Sub:
		return left - right
	case op.Mul:
		return left * right
	case op.Div:
		return left / right
	case op.Mod:
		return left % right
	default:
		panic(errors.InvalidOperation)
	}
}

func evaluateBinaryFloat(left, right float64, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Or:
		return left != 0.0 || right != 0.0
	case op.And:
		return left != 0.0 && right != 0.0
	// Comparision
	case op.Eq:
		return left == right
	case op.Neq:
		return left != right
	case op.Le:
		return left < right
	case op.Ge:
		return left > right
	case op.Lte:
		return left <= right
	case op.Gte:
		return left >= right
	// Arithmetic
	case op.Add:
		return left + right
	case op.Sub:
		return left - right
	case op.Mul:
		return left * right
	case op.Div:
		return left / right
	default:
		panic(errors.InvalidOperation)
	}
}

func evaluateBinaryString(left, right string, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Or:
		return left != "" || right != ""
	case op.And:
		return left != "" && right != ""
	// Comparision
	case op.Eq:
		return left == right
	case op.Neq:
		return left != right
	case op.Le:
		return left < right
	case op.Ge:
		return left > right
	case op.Lte:
		return left <= right
	case op.Gte:
		return left >= right
	// Concatenation
	case op.Add:
		return left + right
	default:
		panic(errors.InvalidOperation)
	}
}

func evaluateBinaryBool(left, right bool, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Or:
		return left || right
	case op.And:
		return left && right
	// Comparision
	case op.Eq:
		return left == right
	case op.Neq:
		return left != right
	default:
		panic(errors.InvalidOperation)
	}
}
