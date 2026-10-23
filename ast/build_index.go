//
// The RXScript author
// built-in functions

package ast

import (
	"strings"
)

type Index struct {
	FuncDef
}

func (p *Index) String() string {
	return "'index'"
}

func NewIndex() *Index {
	return &Index{}
}

func (p *Index) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch v := arg.(type) {
		case *StringDef:
			if len(p.Args) > 1 {
				indexSet := p.Args[1]
				switch index := indexSet.(type) {
				case *StringDef:
					return &IntDef{Value: int64(strings.Index(v.Value, index.Value))}
				default:
					return ThrowSystemErrorWithFormat("RXScript: index second argument must be string, not '%s'", index)
				}
			}
			return arg
		default:
			return ThrowSystemErrorWithFormat("RXScript: index need at least two string argument. not: '%s'", v)
		}
	}
	return &NullDef{}
}
