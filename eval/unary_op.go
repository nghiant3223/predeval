package eval

import (
	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/op"
)

//// Evaluate for op Not Add Sub
func (v *visitor) evaluateUnaryOperator(operandResult interface{}, operator op.Operator) interface{} {
	if operandInt, ok := operandResult.(int64); ok {
		return evaluateUnaryInt(operandInt, operator)
	}

	if operandFloat, ok := operandResult.(float64); ok {
		return evaluateUnaryFloat(operandFloat, operator)
	}

	if operandString, ok := operandResult.(string); ok {
		return evaluateUnaryString(operandString, operator)
	}

	if operandBool, ok := operandResult.(bool); ok {
		return evaluateUnaryBool(operandBool, operator)
	}

	panic(errors.InvalidDataType)
}

func evaluateUnaryInt(operand int64, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Not:
		return operand != 0
	// Arithmetic
	case op.Add:
		return operand
	case op.Sub:
		return -operand
	default:
		panic(errors.InvalidOperation)
	}
}

func evaluateUnaryFloat(operand float64, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Not:
		return operand != 0
	// Arithmetic
	case op.Add:
		return operand
	case op.Sub:
		return -operand
	default:
		panic(errors.InvalidOperation)
	}
}

func evaluateUnaryString(operand string, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Not:
		return operand != ""
	default:
		panic(errors.InvalidOperation)
	}
}

func evaluateUnaryBool(operand bool, operator op.Operator) interface{} {
	switch operator {
	// Logical
	case op.Not:
		return !operand
	default:
		panic(errors.InvalidOperation)
	}
}
