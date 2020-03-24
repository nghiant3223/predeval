package eval

import (
	"testing"

	"github.com/ntncsebku/reloader/ast"
	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/eval"
	"github.com/stretchr/testify/suite"
)

type ParserSuite struct {
	suite.Suite
	parser    ast.Parser
	evaluator eval.Evaluator
}

type TestCase struct {
	Input    string
	Args     map[string]string
	Expected interface{}
}

func TestEvaluatorSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

func (s *ParserSuite) SetupTest() {
	s.parser = ast.NewParser()
	s.evaluator = eval.NewEvaluator()
}

func (s *ParserSuite) TestInvalidOperation() {
	for _, t := range []TestCase{
		{
			Input:    "2 > false",
			Expected: errors.InvalidOperation,
		},
		{
			Input:    "1 and  3 - '.'",
			Expected: errors.InvalidOperation,
		},
		{
			Input:    "3 % '.'",
			Expected: errors.InvalidOperation,
		},
		{
			Input:    "x %  3 - x",
			Expected: errors.InvalidOperation,
			Args:     newVariableMap("x", "'a string'"),
		},
		{
			Input:    "false and true and 'dsf'",
			Expected: errors.InvalidOperation,
		},
		{
			Input:    "+true - false",
			Expected: errors.InvalidOperation,
		},
		{
			Input:    "1 + -'sdf'",
			Expected: errors.InvalidOperation,
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.Equal(t.Expected, err)
		s.Equal(nil, result)
	}
}

func (s *ParserSuite) TestDivideByZero() {
	for _, t := range []TestCase{
		{
			Input:    "100/0",
			Expected: errors.DivideByZero,
		},
		{
			Input:    "100.0/0",
			Expected: errors.DivideByZero,
		},
		{
			Input:    "100.0/0.0",
			Expected: errors.DivideByZero,
		},
		{
			Input:    "100/0.0",
			Expected: errors.DivideByZero,
		},
		{
			Input:    "100/(100*10-10*100)",
			Expected: errors.DivideByZero,
		},
		{
			Input:    "100/((x-1)*(x-1) - x*x + 2*x - 1)",
			Expected: errors.DivideByZero,
			Args:     newVariableMap("x", "1000"),
		},
		{
			Input:    "100/x",
			Expected: errors.DivideByZero,
			Args:     newVariableMap("x", "100"),
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.Equal(t.Expected, err)
		s.Equal(nil, result)
	}
}

func (s *ParserSuite) TestSuccess() {
	for _, t := range []TestCase{
		{
			Input:    "100 ++ x",
			Expected: int64(200),
			Args:     newVariableMap("X", "100"),
		},
		{
			Input:    "100 ++ X",
			Expected: int64(200),
			Args:     newVariableMap("x", "100"),
		},
		{
			Input:    "100 ++ x",
			Expected: int64(200),
			Args:     newVariableMap("x", "100"),
		},
		{
			Input:    "100 > (100-12) and x in [10+23, false, 23]",
			Expected: true,
			Args:     newVariableMap("x", "33"),
		},
		{
			Input:    "100 > (100-12) and x in [10+23, false, 23]",
			Expected: true,
			Args:     newVariableMap("x", "false"),
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.NoError(err)
		s.Equal(t.Expected, result)
	}
}
