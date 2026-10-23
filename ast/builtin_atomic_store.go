//
// The RXScript author
// built-in functions

package ast

type AtomicIntStore struct {
	FuncDef
}

func (p *AtomicIntStore) String() string {
	return "'storeAtomicInt'"
}

func NewAtomicIntStore() *AtomicIntStore {
	return &AtomicIntStore{}
}

func (p *AtomicIntStore) Execute(e RxeHandler) any {
	if len(p.Args) != 2 {
		return ThrowSystemError("RXScript: atomicIntStore: invalid number of arguments")
	}
	ac := p.Args[0]
	vl := p.Args[1]

	switch w := ac.(type) {
	case *AtomicIntDef:
		switch v := vl.(type) {
		case *IntDef:
			w.Value.Store(v.Value)
		default:
			return ThrowSystemErrorWithFormat("RXScript: atomicIntStore: invalid type for AtomicIntDef, not '%s'", v)
		}
		return w
	default:
		return ThrowSystemError("RXScript: atomicIntStore: argument must be atomic value.")
	}
}
