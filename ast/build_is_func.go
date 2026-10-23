//
// The RXScript author
// built-in functions

package ast

type IsFunction struct {
	FuncDef
}

func (p *IsFunction) String() string {
	return "'isFunction'"
}

func NewIsFunction() *IsFunction {
	return &IsFunction{}
}

func (p *IsFunction) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch arg.(type) {
		case *FunctionDef:
			return &BoolDef{Value: true}
		case *ClassBodyDef:
			return &BoolDef{Value: true}
		}
	}
	return &BoolDef{Value: false}
}
