package expression

import (
	"fmt"
	"strconv"
)

// Int

type IntLiteral struct {
	Value int64
}

func NewIntLiteral(str string) *IntLiteral {
	v, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		return &IntLiteral{Value: 0}
	}
	return &IntLiteral{Value: v}
}

func (l *IntLiteral) String() string {
	return fmt.Sprintf("Int(%d)", l.Value)
}

// Bool

type BoolLiteral struct {
	Value bool
}

func NewBoolLiteral(str string) *BoolLiteral {
	v, e := strconv.ParseBool(str)
	if e != nil {
		return &BoolLiteral{Value: false}
	}
	return &BoolLiteral{Value: v}
}

func (l *BoolLiteral) String() string {
	return fmt.Sprintf("Bool(%t)", l.Value)
}

// Float

type FloatLiteral struct {
	Value float64
}

func NewFloatLiteral(str string) *FloatLiteral {
	v, e := strconv.ParseFloat(str, 64)
	if e != nil {
		return &FloatLiteral{Value: 0}
	}
	return &FloatLiteral{Value: v}
}

func (l *FloatLiteral) String() string {
	return fmt.Sprintf("Float(%f)", l.Value)
}

// String

type StringLiteral struct {
	Value string
}

func NewStringLiteral(v string) *StringLiteral {
	realV := v[1 : len(v)-1]
	return &StringLiteral{Value: realV}
}

func (l *StringLiteral) String() string {
	return fmt.Sprintf("String(%s)", l.Value)
}
