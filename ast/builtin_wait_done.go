//
// The RXScript author
// built-in functions

package ast

type WaitDone struct {
	FuncDef
}

func (p *WaitDone) String() string {
	return "'waitDone'"
}

func NewWaitDone() *WaitDone {
	return &WaitDone{}
}

func (p *WaitDone) Execute(e RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: waitDone: invalid number of arguments")
	}
	wg := p.Args[0]

	switch w := wg.(type) {
	case *WaitDef:
		w.Value.Done()
		return w
	default:
		return ThrowSystemError("RXScript: waitDone: argument must be wait resource.")
	}
}
