//
// The RXScript author
// built-in functions

package ast

type IsObject struct {
	FuncDef
}

func (p *IsObject) String() string {
	return "'isObject'"
}

func NewIsObject() *IsObject {
	return &IsObject{}
}

func (p *IsObject) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *Class:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
