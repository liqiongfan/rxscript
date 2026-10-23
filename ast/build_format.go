//
// The RXScript author
// built-in functions

package ast

import "fmt"

type Format struct {
	FuncDef
}

func (p *Format) String() string {
	return "'format'"
}

func NewFormat() *Format {
	return &Format{}
}

// Execute format("%s%s%s", "a", "b", "c")
func (p *Format) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch v := arg.(type) {
		case *StringDef:
			args := make([]any, 0, len(p.Args))
			if len(p.Args) > 1 {
				for _, x := range p.Args[1:] {
					args = append(args, rawValue(x, rxe))
				}
			}
			return &StringDef{Value: fmt.Sprintf(v.Value, args...)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: format need at least one string argument. not: '%s'", v)
		}
	}
	// no argument passed.
	return &NullDef{}
}
