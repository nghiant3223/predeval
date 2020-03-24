package eval

import "github.com/ntncsebku/reloader/ast/expression"

type Visitor interface {
	Visit(c expression.Expression) (interface{}, error)
	visitBinaryExpression(c *expression.BinaryExpression) (interface{}, error)
	visitUnaryExpression(c *expression.UnaryExpression) (interface{}, error)
	visitMembershipExpression(c *expression.MembershipExpression) (interface{}, error)
	visitVariable(c *expression.Variable) (interface{}, error)
	visitIntLiteral(c *expression.IntLiteral) (interface{}, error)
	visitFloatLiteral(c *expression.FloatLiteral) (interface{}, error)
	visitStringLiteral(c *expression.StringLiteral) (interface{}, error)
	visitBoolLiteral(c *expression.BoolLiteral) (interface{}, error)
}
