package expression

import (
	"fmt"
)

type Variable struct {
	Name string
}

func NewVariable(name string) *Variable {
	return &Variable{Name: name}
}

func (e *Variable) String() string {
	return fmt.Sprintf("Var(%s)", e.Name)
}
