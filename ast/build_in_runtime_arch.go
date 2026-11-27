//
// The RXScript author
// built-in functions

package ast

import "runtime"

type RuntimeARCH struct {
	FuncDef
}

func (p *RuntimeARCH) String() string {
	return "'RuntimeARCH'"
}

func NewRuntimeARCH() *RuntimeARCH {
	return &RuntimeARCH{}
}

func (p *RuntimeARCH) Execute(rxe RxeHandler) any {
	return &StringDef{Value: runtime.GOARCH}
}
