// Code generated from Predicate.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Predicate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PredicateListener is a complete listener for a parse tree produced by PredicateParser.
type PredicateListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterExp1 is called when entering the exp1 production.
	EnterExp1(c *Exp1Context)

	// EnterExp2 is called when entering the exp2 production.
	EnterExp2(c *Exp2Context)

	// EnterExp3 is called when entering the exp3 production.
	EnterExp3(c *Exp3Context)

	// EnterExp4 is called when entering the exp4 production.
	EnterExp4(c *Exp4Context)

	// EnterExp5 is called when entering the exp5 production.
	EnterExp5(c *Exp5Context)

	// EnterExp6 is called when entering the exp6 production.
	EnterExp6(c *Exp6Context)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitExp1 is called when exiting the exp1 production.
	ExitExp1(c *Exp1Context)

	// ExitExp2 is called when exiting the exp2 production.
	ExitExp2(c *Exp2Context)

	// ExitExp3 is called when exiting the exp3 production.
	ExitExp3(c *Exp3Context)

	// ExitExp4 is called when exiting the exp4 production.
	ExitExp4(c *Exp4Context)

	// ExitExp5 is called when exiting the exp5 production.
	ExitExp5(c *Exp5Context)

	// ExitExp6 is called when exiting the exp6 production.
	ExitExp6(c *Exp6Context)
}
