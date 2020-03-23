package eval

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Visitor interface {
	VisitBinaryExpression(c antlr.BaseParserRuleContext)
	VisitUnaryExpression(c antlr.BaseParserRuleContext)
	VisitMembershipExpression(c antlr.BaseInterpreterRuleContext)
}
