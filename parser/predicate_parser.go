// Code generated from Predicate.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Predicate

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 135,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 28, 10, 3, 12, 3, 14, 3, 31, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 39, 10, 4, 12, 4, 14, 4, 42, 11, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	7, 5, 71, 10, 5, 12, 5, 14, 5, 74, 11, 5, 5, 5, 76, 10, 5, 3, 5, 7, 5,
	79, 10, 5, 12, 5, 14, 5, 82, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 93, 10, 6, 12, 6, 14, 6, 96, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 110,
	10, 7, 12, 7, 14, 7, 113, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 5, 8, 122, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 5, 9, 133, 10, 9, 3, 9, 2, 7, 4, 6, 8, 10, 12, 10, 2, 4, 6, 8, 10, 12,
	14, 16, 2, 2, 2, 150, 2, 18, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 32, 3, 2,
	2, 2, 8, 43, 3, 2, 2, 2, 10, 83, 3, 2, 2, 2, 12, 97, 3, 2, 2, 2, 14, 121,
	3, 2, 2, 2, 16, 132, 3, 2, 2, 2, 18, 19, 5, 4, 3, 2, 19, 20, 7, 2, 2, 3,
	20, 3, 3, 2, 2, 2, 21, 22, 8, 3, 1, 2, 22, 23, 5, 6, 4, 2, 23, 29, 3, 2,
	2, 2, 24, 25, 12, 4, 2, 2, 25, 26, 7, 9, 2, 2, 26, 28, 5, 6, 4, 2, 27,
	24, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2,
	2, 30, 5, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 33, 8, 4, 1, 2, 33, 34, 5,
	8, 5, 2, 34, 40, 3, 2, 2, 2, 35, 36, 12, 4, 2, 2, 36, 37, 7, 8, 2, 2, 37,
	39, 5, 8, 5, 2, 38, 35, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2,
	2, 40, 41, 3, 2, 2, 2, 41, 7, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 8,
	5, 1, 2, 44, 45, 5, 10, 6, 2, 45, 80, 3, 2, 2, 2, 46, 47, 12, 10, 2, 2,
	47, 48, 7, 20, 2, 2, 48, 79, 5, 10, 6, 2, 49, 50, 12, 9, 2, 2, 50, 51,
	7, 19, 2, 2, 51, 79, 5, 10, 6, 2, 52, 53, 12, 8, 2, 2, 53, 54, 7, 23, 2,
	2, 54, 79, 5, 10, 6, 2, 55, 56, 12, 7, 2, 2, 56, 57, 7, 24, 2, 2, 57, 79,
	5, 10, 6, 2, 58, 59, 12, 6, 2, 2, 59, 60, 7, 22, 2, 2, 60, 79, 5, 10, 6,
	2, 61, 62, 12, 5, 2, 2, 62, 63, 7, 21, 2, 2, 63, 79, 5, 10, 6, 2, 64, 65,
	12, 4, 2, 2, 65, 66, 7, 7, 2, 2, 66, 75, 7, 26, 2, 2, 67, 72, 5, 4, 3,
	2, 68, 69, 7, 30, 2, 2, 69, 71, 5, 4, 3, 2, 70, 68, 3, 2, 2, 2, 71, 74,
	3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 75, 67, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 3,
	2, 2, 2, 77, 79, 7, 27, 2, 2, 78, 46, 3, 2, 2, 2, 78, 49, 3, 2, 2, 2, 78,
	52, 3, 2, 2, 2, 78, 55, 3, 2, 2, 2, 78, 58, 3, 2, 2, 2, 78, 61, 3, 2, 2,
	2, 78, 64, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81,
	3, 2, 2, 2, 81, 9, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 84, 8, 6, 1, 2,
	84, 85, 5, 12, 7, 2, 85, 94, 3, 2, 2, 2, 86, 87, 12, 5, 2, 2, 87, 88, 7,
	16, 2, 2, 88, 93, 5, 12, 7, 2, 89, 90, 12, 4, 2, 2, 90, 91, 7, 17, 2, 2,
	91, 93, 5, 12, 7, 2, 92, 86, 3, 2, 2, 2, 92, 89, 3, 2, 2, 2, 93, 96, 3,
	2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 11, 3, 2, 2, 2, 96,
	94, 3, 2, 2, 2, 97, 98, 8, 7, 1, 2, 98, 99, 5, 14, 8, 2, 99, 111, 3, 2,
	2, 2, 100, 101, 12, 6, 2, 2, 101, 102, 7, 18, 2, 2, 102, 110, 5, 14, 8,
	2, 103, 104, 12, 5, 2, 2, 104, 105, 7, 15, 2, 2, 105, 110, 5, 14, 8, 2,
	106, 107, 12, 4, 2, 2, 107, 108, 7, 14, 2, 2, 108, 110, 5, 14, 8, 2, 109,
	100, 3, 2, 2, 2, 109, 103, 3, 2, 2, 2, 109, 106, 3, 2, 2, 2, 110, 113,
	3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 13, 3, 2,
	2, 2, 113, 111, 3, 2, 2, 2, 114, 115, 7, 16, 2, 2, 115, 122, 5, 14, 8,
	2, 116, 117, 7, 17, 2, 2, 117, 122, 5, 14, 8, 2, 118, 119, 7, 25, 2, 2,
	119, 122, 5, 14, 8, 2, 120, 122, 5, 16, 9, 2, 121, 114, 3, 2, 2, 2, 121,
	116, 3, 2, 2, 2, 121, 118, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 15, 3,
	2, 2, 2, 123, 124, 7, 28, 2, 2, 124, 125, 5, 4, 3, 2, 125, 126, 7, 29,
	2, 2, 126, 133, 3, 2, 2, 2, 127, 133, 7, 3, 2, 2, 128, 133, 7, 4, 2, 2,
	129, 133, 7, 5, 2, 2, 130, 133, 7, 6, 2, 2, 131, 133, 7, 12, 2, 2, 132,
	123, 3, 2, 2, 2, 132, 127, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 132, 129,
	3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133, 17, 3, 2,
	2, 2, 14, 29, 40, 72, 75, 78, 80, 92, 94, 109, 111, 121, 132,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "'*'", "'/'", "'+'", "'-'",
	"'%'", "'!='", "'=='", "'>='", "'<='", "'<'", "'>'", "'!'", "'['", "']'",
	"'('", "')'", "','", "'''", "'\"'", "'.'",
}
var symbolicNames = []string{
	"", "BOOL_LIT", "INT_LIT", "STRING_LIT", "FLOAT_LIT", "IN", "AND", "OR",
	"TRUE", "FALSE", "VAR", "WHITESPACE", "MUL", "DIV", "ADD", "SUB", "MOD",
	"NEQ", "EQ", "GTE", "LTE", "LT", "GT", "NOT", "LB", "RB", "LP", "RP", "CM",
	"SQ", "DQ", "DT", "SD",
}

var ruleNames = []string{
	"start", "exp", "exp1", "exp2", "exp3", "exp4", "exp5", "exp6",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PredicateParser struct {
	*antlr.BaseParser
}

func NewPredicateParser(input antlr.TokenStream) *PredicateParser {
	this := new(PredicateParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Predicate.g4"

	return this
}

// PredicateParser tokens.
const (
	PredicateParserEOF        = antlr.TokenEOF
	PredicateParserBOOL_LIT   = 1
	PredicateParserINT_LIT    = 2
	PredicateParserSTRING_LIT = 3
	PredicateParserFLOAT_LIT  = 4
	PredicateParserIN         = 5
	PredicateParserAND        = 6
	PredicateParserOR         = 7
	PredicateParserTRUE       = 8
	PredicateParserFALSE      = 9
	PredicateParserVAR        = 10
	PredicateParserWHITESPACE = 11
	PredicateParserMUL        = 12
	PredicateParserDIV        = 13
	PredicateParserADD        = 14
	PredicateParserSUB        = 15
	PredicateParserMOD        = 16
	PredicateParserNEQ        = 17
	PredicateParserEQ         = 18
	PredicateParserGTE        = 19
	PredicateParserLTE        = 20
	PredicateParserLT         = 21
	PredicateParserGT         = 22
	PredicateParserNOT        = 23
	PredicateParserLB         = 24
	PredicateParserRB         = 25
	PredicateParserLP         = 26
	PredicateParserRP         = 27
	PredicateParserCM         = 28
	PredicateParserSQ         = 29
	PredicateParserDQ         = 30
	PredicateParserDT         = 31
	PredicateParserSD         = 32
)

// PredicateParser rules.
const (
	PredicateParserRULE_start = 0
	PredicateParserRULE_exp   = 1
	PredicateParserRULE_exp1  = 2
	PredicateParserRULE_exp2  = 3
	PredicateParserRULE_exp3  = 4
	PredicateParserRULE_exp4  = 5
	PredicateParserRULE_exp5  = 6
	PredicateParserRULE_exp6  = 7
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(PredicateParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *PredicateParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PredicateParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.exp(0)
	}
	{
		p.SetState(17)
		p.Match(PredicateParserEOF)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Exp1() IExp1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp1Context)
}

func (s *ExpContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) OR() antlr.TerminalNode {
	return s.GetToken(PredicateParserOR, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *PredicateParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *PredicateParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, PredicateParserRULE_exp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.exp1(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp)
			p.SetState(22)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(23)
				p.Match(PredicateParserOR)
			}
			{
				p.SetState(24)
				p.exp1(0)
			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IExp1Context is an interface to support dynamic dispatch.
type IExp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp1Context differentiates from other interfaces.
	IsExp1Context()
}

type Exp1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp1Context() *Exp1Context {
	var p = new(Exp1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_exp1
	return p
}

func (*Exp1Context) IsExp1Context() {}

func NewExp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp1Context {
	var p = new(Exp1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_exp1

	return p
}

func (s *Exp1Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp1Context) Exp2() IExp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp2Context)
}

func (s *Exp1Context) Exp1() IExp1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp1Context)
}

func (s *Exp1Context) AND() antlr.TerminalNode {
	return s.GetToken(PredicateParserAND, 0)
}

func (s *Exp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterExp1(s)
	}
}

func (s *Exp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitExp1(s)
	}
}

func (p *PredicateParser) Exp1() (localctx IExp1Context) {
	return p.exp1(0)
}

func (p *PredicateParser) exp1(_p int) (localctx IExp1Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExp1Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExp1Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, PredicateParserRULE_exp1, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.exp2(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExp1Context(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp1)
			p.SetState(33)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(34)
				p.Match(PredicateParserAND)
			}
			{
				p.SetState(35)
				p.exp2(0)
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IExp2Context is an interface to support dynamic dispatch.
type IExp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp2Context differentiates from other interfaces.
	IsExp2Context()
}

type Exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp2Context() *Exp2Context {
	var p = new(Exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_exp2
	return p
}

func (*Exp2Context) IsExp2Context() {}

func NewExp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp2Context {
	var p = new(Exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_exp2

	return p
}

func (s *Exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp2Context) Exp3() IExp3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp3Context)
}

func (s *Exp2Context) Exp2() IExp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp2Context)
}

func (s *Exp2Context) EQ() antlr.TerminalNode {
	return s.GetToken(PredicateParserEQ, 0)
}

func (s *Exp2Context) NEQ() antlr.TerminalNode {
	return s.GetToken(PredicateParserNEQ, 0)
}

func (s *Exp2Context) LT() antlr.TerminalNode {
	return s.GetToken(PredicateParserLT, 0)
}

func (s *Exp2Context) GT() antlr.TerminalNode {
	return s.GetToken(PredicateParserGT, 0)
}

func (s *Exp2Context) LTE() antlr.TerminalNode {
	return s.GetToken(PredicateParserLTE, 0)
}

func (s *Exp2Context) GTE() antlr.TerminalNode {
	return s.GetToken(PredicateParserGTE, 0)
}

func (s *Exp2Context) IN() antlr.TerminalNode {
	return s.GetToken(PredicateParserIN, 0)
}

func (s *Exp2Context) LB() antlr.TerminalNode {
	return s.GetToken(PredicateParserLB, 0)
}

func (s *Exp2Context) RB() antlr.TerminalNode {
	return s.GetToken(PredicateParserRB, 0)
}

func (s *Exp2Context) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *Exp2Context) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Exp2Context) AllCM() []antlr.TerminalNode {
	return s.GetTokens(PredicateParserCM)
}

func (s *Exp2Context) CM(i int) antlr.TerminalNode {
	return s.GetToken(PredicateParserCM, i)
}

func (s *Exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterExp2(s)
	}
}

func (s *Exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitExp2(s)
	}
}

func (p *PredicateParser) Exp2() (localctx IExp2Context) {
	return p.exp2(0)
}

func (p *PredicateParser) exp2(_p int) (localctx IExp2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExp2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExp2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, PredicateParserRULE_exp2, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.exp3(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp2)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(45)
					p.Match(PredicateParserEQ)
				}
				{
					p.SetState(46)
					p.exp3(0)
				}

			case 2:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp2)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(48)
					p.Match(PredicateParserNEQ)
				}
				{
					p.SetState(49)
					p.exp3(0)
				}

			case 3:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp2)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(51)
					p.Match(PredicateParserLT)
				}
				{
					p.SetState(52)
					p.exp3(0)
				}

			case 4:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp2)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(54)
					p.Match(PredicateParserGT)
				}
				{
					p.SetState(55)
					p.exp3(0)
				}

			case 5:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp2)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(57)
					p.Match(PredicateParserLTE)
				}
				{
					p.SetState(58)
					p.exp3(0)
				}

			case 6:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp2)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(60)
					p.Match(PredicateParserGTE)
				}
				{
					p.SetState(61)
					p.exp3(0)
				}

			case 7:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp2)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(63)
					p.Match(PredicateParserIN)
				}
				{
					p.SetState(64)
					p.Match(PredicateParserLB)
				}
				p.SetState(73)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PredicateParserBOOL_LIT)|(1<<PredicateParserINT_LIT)|(1<<PredicateParserSTRING_LIT)|(1<<PredicateParserFLOAT_LIT)|(1<<PredicateParserVAR)|(1<<PredicateParserADD)|(1<<PredicateParserSUB)|(1<<PredicateParserNOT)|(1<<PredicateParserLP))) != 0 {
					{
						p.SetState(65)
						p.exp(0)
					}
					p.SetState(70)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == PredicateParserCM {
						{
							p.SetState(66)
							p.Match(PredicateParserCM)
						}
						{
							p.SetState(67)
							p.exp(0)
						}

						p.SetState(72)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(75)
					p.Match(PredicateParserRB)
				}

			}

		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IExp3Context is an interface to support dynamic dispatch.
type IExp3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp3Context differentiates from other interfaces.
	IsExp3Context()
}

type Exp3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp3Context() *Exp3Context {
	var p = new(Exp3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_exp3
	return p
}

func (*Exp3Context) IsExp3Context() {}

func NewExp3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp3Context {
	var p = new(Exp3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_exp3

	return p
}

func (s *Exp3Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp3Context) Exp4() IExp4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp4Context)
}

func (s *Exp3Context) Exp3() IExp3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp3Context)
}

func (s *Exp3Context) ADD() antlr.TerminalNode {
	return s.GetToken(PredicateParserADD, 0)
}

func (s *Exp3Context) SUB() antlr.TerminalNode {
	return s.GetToken(PredicateParserSUB, 0)
}

func (s *Exp3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterExp3(s)
	}
}

func (s *Exp3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitExp3(s)
	}
}

func (p *PredicateParser) Exp3() (localctx IExp3Context) {
	return p.exp3(0)
}

func (p *PredicateParser) exp3(_p int) (localctx IExp3Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExp3Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExp3Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, PredicateParserRULE_exp3, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.exp4(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp3Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp3)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(85)
					p.Match(PredicateParserADD)
				}
				{
					p.SetState(86)
					p.exp4(0)
				}

			case 2:
				localctx = NewExp3Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp3)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(88)
					p.Match(PredicateParserSUB)
				}
				{
					p.SetState(89)
					p.exp4(0)
				}

			}

		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IExp4Context is an interface to support dynamic dispatch.
type IExp4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp4Context differentiates from other interfaces.
	IsExp4Context()
}

type Exp4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp4Context() *Exp4Context {
	var p = new(Exp4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_exp4
	return p
}

func (*Exp4Context) IsExp4Context() {}

func NewExp4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp4Context {
	var p = new(Exp4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_exp4

	return p
}

func (s *Exp4Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp4Context) Exp5() IExp5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp5Context)
}

func (s *Exp4Context) Exp4() IExp4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp4Context)
}

func (s *Exp4Context) MOD() antlr.TerminalNode {
	return s.GetToken(PredicateParserMOD, 0)
}

func (s *Exp4Context) DIV() antlr.TerminalNode {
	return s.GetToken(PredicateParserDIV, 0)
}

func (s *Exp4Context) MUL() antlr.TerminalNode {
	return s.GetToken(PredicateParserMUL, 0)
}

func (s *Exp4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterExp4(s)
	}
}

func (s *Exp4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitExp4(s)
	}
}

func (p *PredicateParser) Exp4() (localctx IExp4Context) {
	return p.exp4(0)
}

func (p *PredicateParser) exp4(_p int) (localctx IExp4Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExp4Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExp4Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, PredicateParserRULE_exp4, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Exp5()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp4Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp4)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(99)
					p.Match(PredicateParserMOD)
				}
				{
					p.SetState(100)
					p.Exp5()
				}

			case 2:
				localctx = NewExp4Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp4)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(102)
					p.Match(PredicateParserDIV)
				}
				{
					p.SetState(103)
					p.Exp5()
				}

			case 3:
				localctx = NewExp4Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PredicateParserRULE_exp4)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(105)
					p.Match(PredicateParserMUL)
				}
				{
					p.SetState(106)
					p.Exp5()
				}

			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IExp5Context is an interface to support dynamic dispatch.
type IExp5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp5Context differentiates from other interfaces.
	IsExp5Context()
}

type Exp5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp5Context() *Exp5Context {
	var p = new(Exp5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_exp5
	return p
}

func (*Exp5Context) IsExp5Context() {}

func NewExp5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp5Context {
	var p = new(Exp5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_exp5

	return p
}

func (s *Exp5Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp5Context) ADD() antlr.TerminalNode {
	return s.GetToken(PredicateParserADD, 0)
}

func (s *Exp5Context) Exp5() IExp5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp5Context)
}

func (s *Exp5Context) SUB() antlr.TerminalNode {
	return s.GetToken(PredicateParserSUB, 0)
}

func (s *Exp5Context) NOT() antlr.TerminalNode {
	return s.GetToken(PredicateParserNOT, 0)
}

func (s *Exp5Context) Exp6() IExp6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp6Context)
}

func (s *Exp5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterExp5(s)
	}
}

func (s *Exp5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitExp5(s)
	}
}

func (p *PredicateParser) Exp5() (localctx IExp5Context) {
	localctx = NewExp5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PredicateParserRULE_exp5)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PredicateParserADD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(PredicateParserADD)
		}
		{
			p.SetState(113)
			p.Exp5()
		}

	case PredicateParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(PredicateParserSUB)
		}
		{
			p.SetState(115)
			p.Exp5()
		}

	case PredicateParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Match(PredicateParserNOT)
		}
		{
			p.SetState(117)
			p.Exp5()
		}

	case PredicateParserBOOL_LIT, PredicateParserINT_LIT, PredicateParserSTRING_LIT, PredicateParserFLOAT_LIT, PredicateParserVAR, PredicateParserLP:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(118)
			p.Exp6()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExp6Context is an interface to support dynamic dispatch.
type IExp6Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp6Context differentiates from other interfaces.
	IsExp6Context()
}

type Exp6Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp6Context() *Exp6Context {
	var p = new(Exp6Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PredicateParserRULE_exp6
	return p
}

func (*Exp6Context) IsExp6Context() {}

func NewExp6Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp6Context {
	var p = new(Exp6Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredicateParserRULE_exp6

	return p
}

func (s *Exp6Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp6Context) LP() antlr.TerminalNode {
	return s.GetToken(PredicateParserLP, 0)
}

func (s *Exp6Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Exp6Context) RP() antlr.TerminalNode {
	return s.GetToken(PredicateParserRP, 0)
}

func (s *Exp6Context) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(PredicateParserBOOL_LIT, 0)
}

func (s *Exp6Context) INT_LIT() antlr.TerminalNode {
	return s.GetToken(PredicateParserINT_LIT, 0)
}

func (s *Exp6Context) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(PredicateParserSTRING_LIT, 0)
}

func (s *Exp6Context) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(PredicateParserFLOAT_LIT, 0)
}

func (s *Exp6Context) VAR() antlr.TerminalNode {
	return s.GetToken(PredicateParserVAR, 0)
}

func (s *Exp6Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp6Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp6Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.EnterExp6(s)
	}
}

func (s *Exp6Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredicateListener); ok {
		listenerT.ExitExp6(s)
	}
}

func (p *PredicateParser) Exp6() (localctx IExp6Context) {
	localctx = NewExp6Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PredicateParserRULE_exp6)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PredicateParserLP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(PredicateParserLP)
		}
		{
			p.SetState(122)
			p.exp(0)
		}
		{
			p.SetState(123)
			p.Match(PredicateParserRP)
		}

	case PredicateParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(PredicateParserBOOL_LIT)
		}

	case PredicateParserINT_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Match(PredicateParserINT_LIT)
		}

	case PredicateParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Match(PredicateParserSTRING_LIT)
		}

	case PredicateParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.Match(PredicateParserFLOAT_LIT)
		}

	case PredicateParserVAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Match(PredicateParserVAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *PredicateParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	case 2:
		var t *Exp1Context = nil
		if localctx != nil {
			t = localctx.(*Exp1Context)
		}
		return p.Exp1_Sempred(t, predIndex)

	case 3:
		var t *Exp2Context = nil
		if localctx != nil {
			t = localctx.(*Exp2Context)
		}
		return p.Exp2_Sempred(t, predIndex)

	case 4:
		var t *Exp3Context = nil
		if localctx != nil {
			t = localctx.(*Exp3Context)
		}
		return p.Exp3_Sempred(t, predIndex)

	case 5:
		var t *Exp4Context = nil
		if localctx != nil {
			t = localctx.(*Exp4Context)
		}
		return p.Exp4_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PredicateParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PredicateParser) Exp1_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PredicateParser) Exp2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PredicateParser) Exp3_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PredicateParser) Exp4_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
