//
// The RXScript author
// built-in functions

package ast

import "reflect"

type TG struct {
	FuncDef
}

func (p *TG) String() string {
	return "'TG'"
}

func NewTG() *TG {
	return &TG{}
}

// Execute TG(httpResponse)
func (p *TG) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: TG need 1 arguments")
	}

	v := rawValue(p.Args[0], rxe)
	return &SystemResource{obj: reflect.ValueOf(v)}
}
