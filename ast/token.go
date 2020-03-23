package ast

import "github.com/antlr/antlr4/runtime/Go/antlr"

func GetTerminalNodeAt(c *antlr.BaseParserRuleContext, index int) antlr.TerminalNode {
	for j := 0; j < len(c.GetChildren()); j++ {
		child := c.GetChild(j)
		if c2, ok := child.(antlr.TerminalNode); ok {
			if index == 0 {
				return c2
			}
			index--
		}
	}
	return nil
}
