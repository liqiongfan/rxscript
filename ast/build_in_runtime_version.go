//
// The RXScript author
// built-in functions

package ast

import "runtime"

type RuntimeVersion struct {
	FuncDef
}

func (p *RuntimeVersion) String() string {
	return "'runtimeVersion'"
}

func NewRuntimeVersion() *RuntimeVersion {
	return &RuntimeVersion{}
}

func (p *RuntimeVersion) Execute(rxe RxeHandler) any {
	return &StringDef{Value: runtime.Version()}
}
