package ast

import "github.com/ntncsebku/reloader/ast/expression"

type Stack interface {
	Push(item expression.Expression)
	Pop() expression.Expression
	Size() int
}

func NewStack() Stack {
	return &stack{}
}

type stack struct {
	data []expression.Expression
}

func (s *stack) Push(item expression.Expression) {
	s.data = append(s.data, item)
}

func (s *stack) Pop() expression.Expression {
	if len(s.data) < 1 {
		panic("stack is empty")
	}

	// Get the last value from the stack.
	result := s.data[len(s.data)-1]
	// Remove the last element from the stack.
	s.data = s.data[:len(s.data)-1]

	return result
}

func (s *stack) Size() int {
	return len(s.data)
}
