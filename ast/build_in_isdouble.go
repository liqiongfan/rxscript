//
// The RXScript author
// built-in functions

package ast

type IsDouble struct {
	FuncDef
}

func (p *IsDouble) String() string {
	return "'isDouble'"
}

func NewIsDouble() *IsDouble {
	return &IsDouble{}
}

func (p *IsDouble) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *DoubleDef, float64, float32:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
