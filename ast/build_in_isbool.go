//
// The RXScript author
// built-in functions

package ast

type IsBool struct {
	FuncDef
}

func (p *IsBool) String() string {
	return "'isBool'"
}

func NewIsBool() *IsBool {
	return &IsBool{}
}

func (p *IsBool) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *BoolDef, bool:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
