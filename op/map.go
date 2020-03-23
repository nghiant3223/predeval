package op

import "strings"

var mapStringToOp = map[string]Operator{
	"or":  Or,
	"and": And,
	"==":  Eq,
	"!=":  Neq,
	"<":   Le,
	"<=":  Lte,
	">":   Ge,
	">=":  Gte,
	"in":  In,
	"+":   Add,
	"-":   Sub,
	"*":   Mul,
	"/":   Div,
	"%":   Mod,
	"!":   Not,
}

var mapOpToString = map[Operator]string{
	Or:  "or",
	And: "and",
	Eq:  "==",
	Neq: "!=",
	Le:  "<",
	Lte: "<=",
	Ge:  ">",
	Gte: ">=",
	In:  "in",
	Add: "+",
	Sub: "-",
	Mul: "*",
	Div: "/",
	Mod: "%",
	Not: "!",
}

func GetOperatorByString(str string) Operator {
	return mapStringToOp[strings.ToLower(str)]
}
