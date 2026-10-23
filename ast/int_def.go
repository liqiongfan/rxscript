//
// The RXScript author
// IntDef represent any integer

package ast

import (
	"strconv"
)

var _ Node = (*IntDef)(nil)

type IntDef struct {
	Value int64
}

func (x *IntDef) Any(handler RxeHandler) any {
	return x.Value
}

func (x *IntDef) Type() CmdType {
	return IntDefAstCmd
}

func (x *IntDef) Execute(e RxeHandler) any {

	return x
}

func (x *IntDef) String() string {
	return "int(" + strconv.FormatInt(x.Value, 10) + ")"
}
