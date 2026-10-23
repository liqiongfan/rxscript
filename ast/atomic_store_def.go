//
// The RXScript author
// AtomicIntDef represent atomic.Int64 object

package ast

import (
	"strconv"
	"sync/atomic"
)

var _ Node = (*AtomicIntDef)(nil)

type AtomicIntDef struct {
	Value *atomic.Int64
}

func (x *AtomicIntDef) Any(handler RxeHandler) any {
	return x.Value.Load()
}

func (x *AtomicIntDef) Type() CmdType {
	return AtomicDefAstCmd
}

func (x *AtomicIntDef) Execute(e RxeHandler) any {
	return x
}

func (x *AtomicIntDef) String() string {
	return "atomic(" + strconv.FormatInt(x.Value.Load(), 10) + ")"
}
