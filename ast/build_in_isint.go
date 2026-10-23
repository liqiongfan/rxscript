//
// The RXScript author
// built-in functions

package ast

type IsInt struct {
	FuncDef
}

func (p *IsInt) String() string {
	return "'isInt'"
}

func (p *IsInt) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func NewIsInt() *IsInt {
	return &IsInt{}
}

func (p *IsInt) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *IntDef, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
