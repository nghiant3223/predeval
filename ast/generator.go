package ast

import (
	"github.com/ntncsebku/reloader/ast/expression"
	"github.com/ntncsebku/reloader/parser"
)

type generator struct {
	*parser.BasePredicateListener
	stack Stack
}

func newGenerator() *generator {
	return &generator{
		stack: NewStack(),
	}
}

func (g *generator) GetResult() expression.Expression {
	return g.stack.Pop()
}

func (g *generator) ExitExp(c *parser.ExpContext) {
	if c.Exp() == nil {
		return
	}

	right, left := g.stack.Pop(), g.stack.Pop()
	g.stack.Push(expression.NewBinaryExpression(c.OR().GetText(), right, left))
}

func (g *generator) ExitExp1(c *parser.Exp1Context) {
	if c.Exp1() == nil {
		return
	}

	right, left := g.stack.Pop(), g.stack.Pop()
	g.stack.Push(expression.NewBinaryExpression(c.AND().GetText(), right, left))
}

func (g *generator) ExitExp2(c *parser.Exp2Context) {
	if c.Exp2() == nil {
		return
	}

	// For operators other than IN
	if c.IN() == nil {
		op := GetTerminalNodeAt(c.BaseParserRuleContext, 0)
		right, left := g.stack.Pop(), g.stack.Pop()
		g.stack.Push(expression.NewBinaryExpression(op.GetText(), right, left))
		return
	}

	// For IN operator only
	var rights []expression.Expression
	for i := 0; i < len(c.AllExp()); i++ {
		leftmostMember := g.stack.Pop()
		// Prepend leftmost member
		rights = append([]expression.Expression{leftmostMember}, rights...)
	}
	left := g.stack.Pop()
	op := c.IN().GetText()
	g.stack.Push(expression.NewMembershipExpression(op, rights, left))
}

func (g *generator) ExitExp3(c *parser.Exp3Context) {
	if c.Exp3() == nil {
		return
	}

	op := GetTerminalNodeAt(c.BaseParserRuleContext, 0).GetText()
	right, left := g.stack.Pop(), g.stack.Pop()
	g.stack.Push(expression.NewBinaryExpression(op, right, left))
}

func (g *generator) ExitExp4(c *parser.Exp4Context) {
	if c.Exp4() == nil {
		return
	}

	op := GetTerminalNodeAt(c.BaseParserRuleContext, 0).GetText()
	right, left := g.stack.Pop(), g.stack.Pop()
	g.stack.Push(expression.NewBinaryExpression(op, right, left))
}

func (g *generator) ExitExp5(c *parser.Exp5Context) {
	if c.Exp5() == nil {
		return
	}

	operator := GetTerminalNodeAt(c.BaseParserRuleContext, 0).GetText()
	operand := g.stack.Pop()
	g.stack.Push(expression.NewUnaryExpression(operator, operand))
}

func (g *generator) ExitExp6(c *parser.Exp6Context) {
	if c.Exp() != nil {
		return
	}

	if c.INT_LIT() != nil {
		g.stack.Push(expression.NewIntLiteral(c.INT_LIT().GetText()))
		return
	}

	if c.BOOL_LIT() != nil {
		literal := expression.NewBoolLiteral(c.BOOL_LIT().GetText())
		g.stack.Push(literal)
		return
	}

	if c.FLOAT_LIT() != nil {
		g.stack.Push(expression.NewFloatLiteral(c.FLOAT_LIT().GetText()))
		return
	}

	if c.STRING_LIT() != nil {
		g.stack.Push(expression.NewStringLiteral(c.STRING_LIT().GetText()))
		return
	}

	if c.VAR() != nil {
		g.stack.Push(expression.NewVariable(c.VAR().GetText()))
		return
	}
}
