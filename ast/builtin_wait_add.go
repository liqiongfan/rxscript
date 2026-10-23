//
// The RXScript author
// built-in functions

package ast

type WaitAdd struct {
	FuncDef
}

func (p *WaitAdd) String() string {
	return "'waitAdd'"
}

func NewWaitAdd() *WaitAdd {
	return &WaitAdd{}
}

func (p *WaitAdd) Execute(e RxeHandler) any {
	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: waitAdd: invalid number of arguments")
	}
	wg := p.Args[0]
	num := p.Args[1]

	switch w := wg.(type) {
	case *WaitDef:
		switch n := num.(type) {
		case *IntDef:
			num = n.Value
			w.Value.Add(int(n.Value))
		default:
			w.Value.Add(1)
		}
		return w
	default:
		return ThrowSystemError("RXScript: waitAdd: argument must be wait resource.")
	}
}
