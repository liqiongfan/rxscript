//
// The RXScript author
// built-in functions

package ast

import (
	"strings"
	"unicode"
)

type TrimLeft struct {
	FuncDef
}

func (p *TrimLeft) String() string {
	return "'trimLeft'"
}

func NewTrimLeft() *TrimLeft {
	return &TrimLeft{}
}

func (p *TrimLeft) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch v := arg.(type) {
		case *StringDef:
			if len(p.Args) > 1 {
				cutSet := p.Args[1]
				switch cut := cutSet.(type) {
				case *StringDef:
					if cut.Value == "" {
						return &StringDef{Value: strings.TrimLeftFunc(v.Value, unicode.IsSpace)}
					}
					return &StringDef{Value: strings.TrimLeft(v.Value, cut.Value)}
				default:
					return ThrowSystemErrorWithFormat("RXScript: trimLeft second argument must be string, not '%s'", cut)
				}
			}
			return &StringDef{Value: strings.TrimLeftFunc(v.Value, unicode.IsSpace)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: trimLeft need at least one string argument. not: '%s'", v)
		}
	}
	// no argument passed.
	return &NullDef{}
}
