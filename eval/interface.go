package eval

import "github.com/ntncsebku/reloader/ast/expression"

type Visitor interface {
	Visit(c expression.Expression) interface{}
	visitBinaryExpression(c *expression.BinaryExpression) interface{}
	visitUnaryExpression(c *expression.UnaryExpression) interface{}
	visitMembershipExpression(c *expression.MembershipExpression) interface{}
	visitVariable(c *expression.Variable) interface{}
	visitIntLiteral(c *expression.IntLiteral) interface{}
	visitFloatLiteral(c *expression.FloatLiteral) interface{}
	visitStringLiteral(c *expression.StringLiteral) interface{}
	visitBoolLiteral(c *expression.BoolLiteral) interface{}
}

func newEvalVisitor(args map[string]string) *visitor {
	return &visitor{mapVariableNameToValue: args}
}
