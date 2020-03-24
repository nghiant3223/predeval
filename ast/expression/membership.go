package expression

import (
	"fmt"
	"strings"

	"github.com/ntncsebku/reloader/op"
)

type MembershipExpression struct {
	Operator     op.Operator
	RightOperand []Expression
	LeftOperand  Expression
}

func NewMembershipExpression(operator string, rightOperand []Expression, leftOperand Expression) *MembershipExpression {
	return &MembershipExpression{
		Operator:     op.GetOperatorByString(operator),
		RightOperand: rightOperand,
		LeftOperand:  leftOperand,
	}
}

func (e *MembershipExpression) String() string {
	var members []string
	for _, member := range e.RightOperand {
		members = append(members, member.String())
	}
	membersString := strings.Join(members, ",")

	return fmt.Sprintf("Membership(%s,%s,[%s])", e.Operator, e.LeftOperand, membersString)
}
