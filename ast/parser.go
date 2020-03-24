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

type predicateParser struct{}

func NewParser() Parser {
	return &predicateParser{}
}

func (p *predicateParser) Parse(input string) (exp expression.Expression, err error) {
	defer func() {
		// If there is any error while lexing or parsing,
		// which is triggered by antlr.ParseTreeWalkerDefault.Walk(astGenerator, antlrParser.Start())
		if r := recover(); r != nil {
			exp = nil
			err, _ = r.(errors.ListenerError)
		}
	}()

	expString := antlr.NewInputStream(input)
	listener := &errors.ErrorListener{}

	// Create the Lexer
	lexer := parser.NewPredicateLexer(expString)
	lexer.AddErrorListener(listener)

	// Create input stream to lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	antlrParser := parser.NewPredicateParser(stream)
	antlrParser.AddErrorListener(listener)

	// Parse the expression
	astGenerator := newGenerator()
	antlr.ParseTreeWalkerDefault.Walk(astGenerator, antlrParser.Start())

	return astGenerator.GetResult(), nil
}
