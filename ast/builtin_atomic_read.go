//
// The RXScript author
// built-in functions

package ast

type AtomicIntLoad struct {
	FuncDef
}

func (p *AtomicIntLoad) String() string {
	return "'readAtomicInt'"
}

func NewAtomicIntLoad() *AtomicIntLoad {
	return &AtomicIntLoad{}
}

func (p *AtomicIntLoad) Execute(e RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: readAtomicInt: invalid number of arguments")
	}
	ac := p.Args[0]

	switch w := ac.(type) {
	case *AtomicIntDef:
		return &IntDef{Value: w.Value.Load()}
	default:
		return ThrowSystemError("RXScript: readAtomicInt: argument must be atomic value.")
	}
}
