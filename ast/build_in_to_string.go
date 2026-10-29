//
// The RXScript author
// built-in functions

package ast

import (
	"fmt"
)

type ToString struct {
	FuncDef
}

func (p *ToString) String() string {
	return "'toString'"
}

func NewToString() *ToString {
	return &ToString{}
}

func (p *ToString) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: toString: need at least 1 argument.")
	}

	v := p.Args[0]
	switch vv := v.(type) {
	case *DoubleDef:
		return &StringDef{Value: fmt.Sprintf(`%g`, vv.Value)}
	default:
		return &StringDef{Value: fmt.Sprintf(`%v`, vv)}
	}
}
