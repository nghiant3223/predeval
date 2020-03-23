package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ntncsebku/reloader/ast/expression"
	"github.com/ntncsebku/reloader/errors"
	"github.com/ntncsebku/reloader/parser"
)

type Parser interface {
	Parse(input string) (expression.Expression, error)
}

type predicateParser struct {
}

func NewParser() Parser {
	return &predicateParser{}
}

func (p *predicateParser) Parse(input string) (expression.Expression, error) {
	exp := antlr.NewInputStream(input)
	listener := &errors.ErrorListener{}

	// Create the Lexer
	lexer := parser.NewPredicateLexer(exp)
	lexer.AddErrorListener(listener)

	// Create input stream to lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	antlrParser := parser.NewPredicateParser(stream)
	antlrParser.AddErrorListener(listener)

	// Parse the exp
	condExprListener := newGenerator()
	antlr.ParseTreeWalkerDefault.Walk(condExprListener, antlrParser.Start())

	// Check for any errors occur while lexing & parsing
	if listener.ErrorCount > 0 {
		return nil, errors.SyntaxError
	}

	return condExprListener.GetResult(), nil
}
