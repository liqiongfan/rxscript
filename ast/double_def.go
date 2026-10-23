//
// The RXScript author
// DoubleDef represent float value

package ast

import (
	"strconv"
)

var _ Node = (*DoubleDef)(nil)

type DoubleDef struct {
	Value float64
}

func (x *DoubleDef) Any(handler RxeHandler) any {
	return x.Value
}

func (x *DoubleDef) Type() CmdType {
	return DoubleDefAstCmd
}

func (x *DoubleDef) Execute(e RxeHandler) any {
	return x
}

func (x *DoubleDef) String() string {
	return "double(" + strconv.FormatFloat(x.Value, 'f', -1, 64) + ")"
}
