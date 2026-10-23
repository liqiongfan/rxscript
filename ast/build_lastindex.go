//
// The RXScript author
// built-in functions

package ast

import (
	"strings"
)

type LastIndex struct {
	FuncDef
}

func (p *LastIndex) String() string {
	return "'lastIndex'"
}

func NewLastIndex() *LastIndex {
	return &LastIndex{}
}

func (p *LastIndex) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch v := arg.(type) {
		case *StringDef:
			if len(p.Args) > 1 {
				indexSet := p.Args[1]
				switch index := indexSet.(type) {
				case *StringDef:
					return &IntDef{Value: int64(strings.LastIndex(v.Value, index.Value))}
				default:
					return ThrowSystemErrorWithFormat("RXScript: lastIndex second argument must be string, not '%s'", index)
				}
			}
			return arg
		default:
			return ThrowSystemErrorWithFormat("RXScript: lastIndex need at least two string argument. not: '%s'", v)
		}
	}
	// no argument passed.
	return &NullDef{}
}
