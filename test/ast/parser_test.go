package ast

import (
	"testing"

	"github.com/ntncsebku/reloader/ast"
	"github.com/ntncsebku/reloader/errors"
	"github.com/stretchr/testify/suite"
)

type ParserSuite struct {
	suite.Suite
	parser ast.Parser
}

type TestCase struct {
	Input    string
	Expected interface{}
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

func (s *ParserSuite) SetupTest() {
	s.parser = ast.NewParser()
}

func (s *ParserSuite) TestSyntaxError() {
	for _, t := range []TestCase{
		{
			Input:    "$vd > 10",
			Expected: errors.SyntaxError,
		},
		{
			Input:    "3x > 10",
			Expected: errors.SyntaxError,
		},
		{
			Input:    "x > 10 + ",
			Expected: errors.SyntaxError,
		},
		{
			Input:    "x and and x",
			Expected: errors.SyntaxError,
		},
		{
			Input:    "[1,3] > 0",
			Expected: errors.SyntaxError,
		},
		{
			Input:    "x in ([1,2])",
			Expected: errors.SyntaxError,
		},
	} {
		result, err := s.parser.Parse(t.Input)
		s.Equal(t.Expected, err)
		s.Equal(nil, result)
	}
}

func (s *ParserSuite) TestBasic() {
	for _, t := range []TestCase{
		{
			Input:    "x",
			Expected: "Var(x)",
		},
		{
			Input:    "false",
			Expected: "Bool(false)",
		},
		{
			Input:    "3",
			Expected: "Int(3)",
		},
		{
			Input:    "'xfd'",
			Expected: "String(xfd)",
		},
		{
			Input:    "\"xfd\"",
			Expected: "String(xfd)",
		},
	} {
		result, err := s.parser.Parse(t.Input)
		s.NoError(err)
		resultString := result.String()
		s.Equal(t.Expected, resultString)
	}
}

func (s *ParserSuite) TestCaseInsensitive() {
	for _, t := range []TestCase{
		{
			Input:    "Xx aND 10",
			Expected: "Binary(and,Var(Xx),Int(10))",
		},
		{
			Input:    "Xx In [10]",
			Expected: "Membership(in,Var(Xx),[Int(10)])",
		},
	} {
		result, err := s.parser.Parse(t.Input)
		s.NoError(err)
		resultString := result.String()
		s.Equal(t.Expected, resultString)
	}
}

func (s *ParserSuite) TestIn() {
	for _, t := range []TestCase{
		{
			Input:    "Xx In [10]",
			Expected: "Membership(in,Var(Xx),[Int(10)])",
		},
		{
			Input:    "Xx In []",
			Expected: "Membership(in,Var(Xx),[])",
		},
		{
			Input:    "Xx In [1, y in [2,z+3]]",
			Expected: "Membership(in,Var(Xx),[Int(1),Membership(in,Var(y),[Int(2),Binary(+,Var(z),Int(3))])])",
		},
	} {
		result, err := s.parser.Parse(t.Input)
		s.NoError(err)
		resultString := result.String()
		s.Equal(t.Expected, resultString)
	}
}

func (s *ParserSuite) TestComparision() {
	for _, t := range []TestCase{
		{
			Input:    "x_2 > 10",
			Expected: "Binary(>,Var(x_2),Int(10))",
		},
		{
			Input:    "x > +10",
			Expected: "Binary(>,Var(x),Unary(+,Int(10)))",
		},
		{
			Input:    "!x > +10",
			Expected: "Binary(>,Unary(!,Var(x)),Unary(+,Int(10)))",
		},
		{
			Input:    "!x > -10.02",
			Expected: "Binary(>,Unary(!,Var(x)),Unary(-,Float(10.020000)))",
		},
		{
			Input:    "x and 2 or y",
			Expected: "Binary(or,Binary(and,Var(x),Int(2)),Var(y))",
		},
		{
			Input:    "x and (2 or y)",
			Expected: "Binary(and,Var(x),Binary(or,Int(2),Var(y)))",
		},
		{
			Input:    "x and (2 or !y)",
			Expected: "Binary(and,Var(x),Binary(or,Int(2),Unary(!,Var(y))))",
		},
		{
			Input:    "x and (2 or !y)",
			Expected: "Binary(and,Var(x),Binary(or,Int(2),Unary(!,Var(y))))",
		},
		{
			Input:    "x and (1+2 or !y)",
			Expected: "Binary(and,Var(x),Binary(or,Binary(+,Int(1),Int(2)),Unary(!,Var(y))))",
		},
	} {
		result, err := s.parser.Parse(t.Input)
		s.NoError(err)
		resultString := result.String()
		s.Equal(t.Expected, resultString)
	}
}

func (s *ParserSuite) TestArithmetic() {
	for _, t := range []TestCase{
		{
			Input:    "(1) + 2",
			Expected: "Binary(+,Int(1),Int(2))",
		},
		{
			Input:    "((x) + 1) + +(2)",
			Expected: "Binary(+,Binary(+,Var(x),Int(1)),Unary(+,Int(2)))",
		},
		{
			Input:    "(1+2)*x + y",
			Expected: "Binary(+,Binary(*,Binary(+,Int(1),Int(2)),Var(x)),Var(y))",
		},
		{
			Input:    "x + y * 2",
			Expected: "Binary(+,Var(x),Binary(*,Var(y),Int(2)))",
		},
		{
			Input:    "x + y * 2 - 6%4",
			Expected: "Binary(-,Binary(+,Var(x),Binary(*,Var(y),Int(2))),Binary(%,Int(6),Int(4)))",
		},
		{
			Input:    "x + y * 2 - -6%4",
			Expected: "Binary(-,Binary(+,Var(x),Binary(*,Var(y),Int(2))),Binary(%,Unary(-,Int(6)),Int(4)))",
		},
		{
			Input:    "x + + x",
			Expected: "Binary(+,Var(x),Unary(+,Var(x)))",
		},
		{
			Input:    "x + + + x",
			Expected: "Binary(+,Var(x),Unary(+,Unary(+,Var(x))))",
		},
	} {
		result, err := s.parser.Parse(t.Input)
		s.NoError(err)
		resultString := result.String()
		s.Equal(t.Expected, resultString)
	}
}

func (s *ParserSuite) TestPrecedence() {
	for _, t := range []TestCase{
		{
			Input:    "!x % y + 3 > 2 and 0 oR 1",
			Expected: "Binary(or,Binary(and,Binary(>,Binary(+,Binary(%,Unary(!,Var(x)),Var(y)),Int(3)),Int(2)),Int(0)),Int(1))",
		},
		{
			Input:    "!x % +y + 3 > 2 and 0 or 1",
			Expected: "Binary(or,Binary(and,Binary(>,Binary(+,Binary(%,Unary(!,Var(x)),Unary(+,Var(y))),Int(3)),Int(2)),Int(0)),Int(1))",
		},
		{
			Input:    "x > y % z in [1,2,3]",
			Expected: "Membership(in,Binary(>,Var(x),Binary(%,Var(y),Var(z))),[Int(1),Int(2),Int(3)])",
		},
		{
			Input:    "x > y % z in [1+'2',x>3]",
			Expected: "Membership(in,Binary(>,Var(x),Binary(%,Var(y),Var(z))),[Binary(+,Int(1),String(2)),Binary(>,Var(x),Int(3))])",
		},
		{
			Input:    "!x in [1+'2',x>3]",
			Expected: "Membership(in,Unary(!,Var(x)),[Binary(+,Int(1),String(2)),Binary(>,Var(x),Int(3))])",
		},
		{
			Input:    "!x and !y in [1+'2',x>3]",
			Expected: "Binary(and,Unary(!,Var(x)),Membership(in,Unary(!,Var(y)),[Binary(+,Int(1),String(2)),Binary(>,Var(x),Int(3))]))",
		},
	} {
		result, err := s.parser.Parse(t.Input)
		s.NoError(err)
		resultString := result.String()
		s.Equal(t.Expected, resultString)
	}
}
