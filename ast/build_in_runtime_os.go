//
// The RXScript author
// built-in functions

package ast

import "runtime"

type RuntimeOS struct {
	FuncDef
}

func (p *RuntimeOS) String() string {
	return "'runtimeOS'"
}

func NewRuntimeOS() *RuntimeOS {
	return &RuntimeOS{}
}

func (p *RuntimeOS) Execute(rxe RxeHandler) any {
	return &StringDef{Value: runtime.GOOS}
}
