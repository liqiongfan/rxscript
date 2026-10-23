//
// The RXScript author
// built-in functions

package ast

type IsString struct {
	FuncDef
}

func (p *IsString) String() string {
	return "'isString'"
}

func NewIsString() *IsString {
	return &IsString{}
}

func (p *IsString) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *StringDef, string:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
