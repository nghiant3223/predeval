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
			Input:    "(false or true) and 'dsf'",
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

func (s *ParserSuite) TestArithmetic() {
	for _, t := range []TestCase{
		{
			Input:    "100 ++ x",
			Expected: int64(200),
			Args:     newVariableMap("X", "100"),
		},
		{
			Input:    "x/2",
			Expected: int64(50),
			Args:     newVariableMap("X", "100"),
		},
		{
			Input:    "x/2.0",
			Expected: 50.0,
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
			Input:    "100 -- x",
			Expected: int64(200),
			Args:     newVariableMap("x", "100"),
		},
		{
			Input:    "1 + 100*20 - 23 * 20 + 32/8 - 32 % 3 ",
			Expected: int64(1543),
		},
		{
			Input:    "100 -- x",
			Expected: int64(0),
			Args:     newVariableMap("x", "-100"),
		},
		{
			Input:    "100 -- x",
			Expected: int64(0),
			Args:     newVariableMap("x", "-100"),
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.NoError(err)
		s.Equal(t.Expected, result)
	}
}

func (s *ParserSuite) TestMembership() {
	for _, t := range []TestCase{
		{
			Input:    "x in [10+23, false, 23]",
			Expected: true,
			Args:     newVariableMap("x", "33"),
		},
		{
			Input:    "x in [10+23, false, 23]",
			Expected: true,
			Args:     newVariableMap("x", "false"),
		},
		{
			Input:    "x in [10+23, false, 23]",
			Expected: true,
			Args:     newVariableMap("x", "false"),
		},
		{
			Input:    "x in [x]",
			Expected: true,
			Args:     newVariableMap("x", "true"),
		},
		{
			Input:    "x in [0, 0, 100]",
			Expected: true,
			Args:     newVariableMap("x", "100"),
		},
		{
			Input:    "x in [0, 0, 100.0]",
			Expected: true,
			Args:     newVariableMap("x", "100"),
		},
		{
			Input:    "x in [0, 0, 100]",
			Expected: true,
			Args:     newVariableMap("x", "100.0"),
		},
		{
			Input:    "x in []",
			Expected: false,
			Args:     newVariableMap("x", "'string'"),
		},
		{
			Input:    "x in [false, 23, 100 in [50+49, 50+50], false]",
			Expected: true,
			Args:     newVariableMap("x", "true"),
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.NoError(err)
		s.Equal(t.Expected, result)
	}
}

func (s *ParserSuite) TestLogical() {
	for _, t := range []TestCase{
		{
			Input:    "!x",
			Expected: false,
			Args:     newVariableMap("x", "150"),
		},
		{
			Input:    "!0",
			Expected: true,
		},
		{
			Input:    "!0.0",
			Expected: true,
		},
		{
			Input:    "!120.0",
			Expected: false,
		},
		{
			Input:    "!false",
			Expected: true,
		},
		{
			Input:    "!true",
			Expected: false,
		},
		{
			Input:    "!''",
			Expected: true,
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.NoError(err)
		s.Equal(t.Expected, result)
	}
}

func (s *ParserSuite) TestComparision() {
	for _, t := range []TestCase{
		{
			Input:    "x > 100 and x < 200 or x % 2 == 0",
			Expected: true,
			Args:     newVariableMap("x", "150"),
		},
		{
			Input:    "x > 100 and x < 200 or x % 2 == 0",
			Expected: true,
			Args:     newVariableMap("x", "50"),
		},
		{
			Input:    "x > 100 and x < 200 or x % 2 == 0",
			Expected: false,
			Args:     newVariableMap("x", "99"),
		},
		{
			Input:    "x >= 100 and (x < 200 or x % 2 == 0)",
			Expected: false,
			Args:     newVariableMap("x", "201"),
		},
		{
			Input:    "x >= 100 and !(x < 200 or x % 2 == 0)",
			Expected: true,
			Args:     newVariableMap("x", "201"),
		},
		{
			Input:    "x >= 100 and !(x < 200 or x % 2 == 0)",
			Expected: true,
			Args:     newVariableMap("x", "201"),
		},
		{
			Input:    "x >= 100 and !(x < 200 or x % 2 == 0)",
			Expected: true,
			Args:     newVariableMap("x", "201"),
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.NoError(err)
		s.Equal(t.Expected, result)
	}
}

func (s *ParserSuite) TestMix() {
	for _, t := range []TestCase{
		{
			Input:    "x < 100 or x > 200 or x % 2 != 0 or x in [false, true, 100 + 50, x > 100]",
			Expected: true,
			Args:     newVariableMap("x", "150"),
		},
		{
			Input:    "(x > 100 and x < 200) and (x % 2 != 0 or x in [false, true, 100 + 50, x > 100])",
			Expected: true,
			Args:     newVariableMap("x", "150"),
		},
	} {
		tree, err := s.parser.Parse(t.Input)
		s.NoError(err)

		result, err := s.evaluator.Evaluate(tree, t.Args)
		s.NoError(err)
		s.Equal(t.Expected, result)
	}
}