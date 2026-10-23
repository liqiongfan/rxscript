//
// The RXScript author
// built-in functions

package ast

type IsArray struct {
	FuncDef
}

func (p *IsArray) String() string {
	return "'isArray'"
}

func NewIsArray() *IsArray {
	return &IsArray{}
}

func (p *IsArray) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *ArrayDef:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
