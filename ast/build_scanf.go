//
// The RXScript author
// built-in functions

package ast

import "fmt"

type Scanf struct {
	FuncDef
}

func (p *Scanf) String() string {
	return "'scanf'"
}

func NewScanf() *Scanf {
	return &Scanf{}
}

// Execute scanf("xxxxx", "%s%s%s", "a", "b", "c")
func (p *Scanf) Execute(rxe RxeHandler) any {
	if len(p.Args) < 3 {
		return ThrowSystemError("RXScript: scanf needs at least 3 arguments")
	}

	str := p.Args[0]
	format := p.Args[1]

	switch d := str.(type) {
	case *StringDef:
		switch f := format.(type) {
		case *StringDef:
			values := make([]any, 0, len(p.Args[2:]))
			for _, arg := range p.Args[2:] {
				values = append(values, rawValueToPtr(arg, rxe))
			}

			_, err := fmt.Sscanf(d.Value, f.Value, values...)
			if err != nil {
				return ThrowSystemError(err)
			}
			return &NullDef{}
		default:
			return ThrowSystemErrorWithFormat("RXScript: scanf second argument need string, not : '%s'", f)
		}
	default:
		return ThrowSystemErrorWithFormat("RXScript: scanf first argument need string, not : '%s'", d)
	}
}
