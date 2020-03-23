package expression

import (
	"fmt"

	"github.com/ntncsebku/reloader/eval"
)

type Variable struct {
	Name  string
	Value interface{}
}

func NewVariable(name string) *Variable {
	return &Variable{Name: name}
}

func (e *Variable) String() string {
	return fmt.Sprintf("Var(%s)", e.Name)
}

func (e *Variable) Accept(v eval.Visitor) {

}
