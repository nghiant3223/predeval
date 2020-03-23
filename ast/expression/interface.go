package expression

import (
	"github.com/ntncsebku/reloader/eval"
)

type Expression interface {
	Accept(v eval.Visitor)
	String() string
}
