//
// The RXScript author
// built-in functions

package ast

type WaitOk struct {
	FuncDef
}

func (p *WaitOk) String() string {
	return "'waitOk'"
}

func NewWaitOk() *WaitOk {
	return &WaitOk{}
}

func (p *WaitOk) Execute(e RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: waitOk: invalid number of arguments")
	}
	wg := p.Args[0]

	switch w := wg.(type) {
	case *WaitDef:
		w.Value.Wait()
		return w
	default:
		return ThrowSystemError("RXScript: waitOk: argument must be wait resource.")
	}
}
