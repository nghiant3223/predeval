package op

type Operator int

const (
	Or Operator = iota + 1
	And
	Eq
	Neq
	Le
	Lte
	Ge
	Gte
	In
	Div
	Mul
	Mod
	Add
	Sub
	Not
)

func (op Operator) String() string {
	return mapOpToString[op]
}
