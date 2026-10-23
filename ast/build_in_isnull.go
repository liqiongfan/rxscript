//
// The RXScript author
// built-in functions

package ast

type IsNull struct {
	FuncDef
}

func (p *IsNull) String() string {
	return "'isNull'"
}

func NewIsNull() *IsNull {
	return &IsNull{}
}

func (p *IsNull) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *NullDef, nil:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
