//
// The RXScript author
// built-in functions

package ast

import "sync/atomic"

type AtomicInt struct {
	FuncDef
}

func (p *AtomicInt) String() string {
	return "'atomic.int'"
}

func NewAtomicInt() *AtomicInt {
	return &AtomicInt{}
}

func (p *AtomicInt) Execute(e RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: atomicInt: invalid number of arguments")
	}
	wg := p.Args[0]
	switch w := wg.(type) {
	case *IntDef:
		v := &atomic.Int64{}
		v.Store(w.Value)
		return &AtomicIntDef{Value: v}
	default:
		return ThrowSystemError("RXScript: atomicInt: argument must be int value.")
	}
}
