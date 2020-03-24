package eval

import (
	"github.com/ntncsebku/reloader/ast/expression"
	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/op"
)

type visitor struct {
	mapVariableNameToValue map[string]string
}

// Mimic method overloading
func (v *visitor) Visit(c expression.Expression) interface{} {
	switch c.(type) {
	case *expression.BinaryExpression:
		exp, _ := c.(*expression.BinaryExpression)
		return v.visitBinaryExpression(exp)
	case *expression.UnaryExpression:
		exp, _ := c.(*expression.UnaryExpression)
		return v.visitUnaryExpression(exp)
	case *expression.MembershipExpression:
		exp, _ := c.(*expression.MembershipExpression)
		return v.visitMembershipExpression(exp)
	case *expression.Variable:
		exp, _ := c.(*expression.Variable)
		return v.visitVariable(exp)
	case *expression.IntLiteral:
		exp, _ := c.(*expression.IntLiteral)
		return v.visitIntLiteral(exp)
	case *expression.StringLiteral:
		exp, _ := c.(*expression.StringLiteral)
		return v.visitStringLiteral(exp)
	case *expression.FloatLiteral:
		exp, _ := c.(*expression.FloatLiteral)
		return v.visitFloatLiteral(exp)
	case *expression.BoolLiteral:
		exp, _ := c.(*expression.BoolLiteral)
		return v.visitBoolLiteral(exp)
	}
	panic(errors.InvalidOperation)
}

func (v *visitor) visitBinaryExpression(c *expression.BinaryExpression) interface{} {
	// Evaluate left operand
	leftResult := v.Visit(c.LeftOperand)

	// If operator is OR and the result of left operand is of bool type and equals TRUE,
	// return TRUE immediately due to short-circuit characteristic
	if leftBool, ok := leftResult.(bool); ok {
		if leftBool && c.Operator == op.Or {
			return true
		}
	}

	// If operator is AND and the result of left operand is of bool type and equals FALSE,
	// return FALSE immediately due to short-circuit characteristic
	if leftBool, ok := leftResult.(bool); ok {
		if leftBool && c.Operator == op.And {
			return false
		}
	}

	// Evaluate right operand
	rightResult := v.Visit(c.RightOperand)

	// Evaluate final result
	return EvaluateBinaryOperator(leftResult, rightResult, c.Operator)
}

func (v *visitor) visitUnaryExpression(c *expression.UnaryExpression) interface{} {
	return v.Visit(c)
}

func (v *visitor) visitMembershipExpression(c *expression.MembershipExpression) interface{} {
	panic("implement me")
	return false
}

func (v *visitor) visitVariable(c *expression.Variable) interface{} {
	panic("implement me")
	return false
}

func (v *visitor) visitIntLiteral(c *expression.IntLiteral) interface{} {
	return c.Value
}

func (v *visitor) visitFloatLiteral(c *expression.FloatLiteral) interface{} {
	return c.Value
}

func (v *visitor) visitStringLiteral(c *expression.StringLiteral) interface{} {
	return c.Value
}

func (v *visitor) visitBoolLiteral(c *expression.BoolLiteral) interface{} {
	return c.Value
}
