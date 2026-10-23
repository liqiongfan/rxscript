//
// The RXScript author
// WaitDef represent Go's WaitGroup

package ast

import "sync"

var _ Node = (*WaitDef)(nil)

type WaitDef struct {
	Value *sync.WaitGroup
}

func (x *WaitDef) Any(handler RxeHandler) any {
	return nil
}

func (x *WaitDef) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (x *WaitDef) Type() CmdType {
	return WaitDefAstCmd
}

func (x *WaitDef) Execute(e RxeHandler) any {
	return x
}

func (x *WaitDef) String() string {
	return "'wait'"
}
