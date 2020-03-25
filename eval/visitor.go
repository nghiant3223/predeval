package eval

import "C"
import (
	"strings"

	"github.com/ntncsebku/reloader/ast/expression"
	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/op"
)

type visitor struct {
	mapVariableNameToValue map[string]string
}

func newEvalVisitor(args map[string]string) *visitor {
	// Lowercase all variables' name
	for key, value := range args {
		newKey := strings.ToLower(key)
		if newKey != key {
			args[newKey] = value
			delete(args, key)
		}
	}

	return &visitor{mapVariableNameToValue: args}
}

// Mimic method overloading
func (v *visitor) Visit(c expression.Expression) (interface{}, error) {
	switch c.(type) {
	case *expression.BinaryExpression:
		exp := c.(*expression.BinaryExpression)
		return v.visitBinaryExpression(exp)
	case *expression.UnaryExpression:
		exp := c.(*expression.UnaryExpression)
		return v.visitUnaryExpression(exp)
	case *expression.MembershipExpression:
		exp := c.(*expression.MembershipExpression)
		return v.visitMembershipExpression(exp)
	case *expression.Variable:
		exp := c.(*expression.Variable)
		return v.visitVariable(exp)
	case *expression.IntLiteral:
		exp := c.(*expression.IntLiteral)
		return v.visitIntLiteral(exp)
	case *expression.StringLiteral:
		exp := c.(*expression.StringLiteral)
		return v.visitStringLiteral(exp)
	case *expression.FloatLiteral:
		exp := c.(*expression.FloatLiteral)
		return v.visitFloatLiteral(exp)
	case *expression.BoolLiteral:
		exp := c.(*expression.BoolLiteral)
		return v.visitBoolLiteral(exp)
	default:
		return nil, errors.InvalidOperation
	}
}

func (v *visitor) visitBinaryExpression(c *expression.BinaryExpression) (interface{}, error) {
	// Evaluate left operand
	leftResult, err := v.Visit(c.LeftOperand)
	if err != nil {
		return nil, err
	}

	// If operator is OR and the result of left operand is not zero value,
	// return TRUE immediately due to short-circuit characteristic
	if c.Operator == op.Or {
		if isZeroValue, err := v.isZeroValue(leftResult); err != nil {
			return nil, err
		} else if !isZeroValue {
			return true, nil
		}
	}

	// If operator is AND and the result of left operand is zero value,
	// return FALSE immediately due to short-circuit characteristic
	if c.Operator == op.And {
		if isZeroValue, err := v.isZeroValue(leftResult); err != nil {
			return nil, err
		} else if isZeroValue {
			return false, nil
		}
	}

	// Evaluate right operand
	rightResult, err := v.Visit(c.RightOperand)
	if err != nil {
		return nil, err
	}

	// Evaluate final result
	return v.evaluateBinaryOperator(leftResult, rightResult, c.Operator)
}

func (v *visitor) visitUnaryExpression(c *expression.UnaryExpression) (interface{}, error) {
	operandResult, err := v.Visit(c.Operand)
	if err != nil {
		return nil, err
	}
	return v.evaluateUnaryOperator(operandResult, c.Operator)
}

func (v *visitor) visitMembershipExpression(c *expression.MembershipExpression) (interface{}, error) {
	// Evaluate left operand
	leftResult, err := v.Visit(c.LeftOperand)
	if err != nil {
		return nil, err
	}

	// evaluateMembershipOperator must be passed a []Expression to implement lazy evaluation
	return v.evaluateMembershipOperator(leftResult, c.RightOperand, c.Operator)
}

func (v *visitor) visitVariable(c *expression.Variable) (interface{}, error) {
	return v.evaluateVariable(c.Name)
}

func (v *visitor) visitIntLiteral(c *expression.IntLiteral) (interface{}, error) {
	return c.Value, nil
}

func (v *visitor) visitFloatLiteral(c *expression.FloatLiteral) (interface{}, error) {
	return c.Value, nil
}

func (v *visitor) visitStringLiteral(c *expression.StringLiteral) (interface{}, error) {
	return c.Value, nil
}

func (v *visitor) visitBoolLiteral(c *expression.BoolLiteral) (interface{}, error) {
	return c.Value, nil
}
