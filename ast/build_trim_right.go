//
// The RXScript author
// built-in functions

package ast

import (
	"strings"
	"unicode"
)

type TrimRight struct {
	FuncDef
}

func (p *TrimRight) String() string {
	return "'trimRight'"
}

func NewTrimRight() *TrimRight {
	return &TrimRight{}
}

func (p *TrimRight) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch v := arg.(type) {
		case *StringDef:
			if len(p.Args) > 1 {
				cutSet := p.Args[1]
				switch cut := cutSet.(type) {
				case *StringDef:
					if cut.Value == "" {
						return &StringDef{Value: strings.TrimRightFunc(v.Value, unicode.IsSpace)}
					}
					return &StringDef{Value: strings.TrimRight(v.Value, cut.Value)}
				default:
					return ThrowSystemErrorWithFormat("RXScript: trimRight second argument must be string, not '%s'", cut)
				}
			}
			return &StringDef{Value: strings.TrimRightFunc(v.Value, unicode.IsSpace)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: trimRight need at least one string argument. not: '%s'", v)
		}
	}
	// no argument passed.
	return &NullDef{}
}
