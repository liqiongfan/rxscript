//
// The RXScript author
// built-in functions

package ast

import "strings"

type Trim struct {
	FuncDef
}

func (p *Trim) String() string {
	return "'trim'"
}

func NewTrim() *Trim {
	return &Trim{}
}

func (p *Trim) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch v := arg.(type) {
		case *StringDef:
			if len(p.Args) > 1 {
				cutSet := p.Args[1]
				switch cut := cutSet.(type) {
				case *StringDef:
					if cut.Value == "" {
						return &StringDef{Value: strings.TrimSpace(v.Value)}
					}
					return &StringDef{Value: strings.Trim(v.Value, cut.Value)}
				default:
					return ThrowSystemErrorWithFormat("RXScript: trim second argument must be string, not '%s'", cut)
				}
			}
			return &StringDef{Value: strings.TrimSpace(v.Value)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: trim need at least one string argument. not: '%s'", v)
		}
	}
	// no argument passed.
	return &NullDef{}
}
