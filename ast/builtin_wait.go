//
// The RXScript author
// built-in functions

package ast

import "sync"

type Wait struct {
	FuncDef
}

func (p *Wait) String() string {
	return "'wait'"
}

func NewWait() *Wait {
	return &Wait{}
}

func (p *Wait) Execute(e RxeHandler) any {
	return &WaitDef{Value: &sync.WaitGroup{}}
}
