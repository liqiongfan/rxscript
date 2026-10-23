//
// The RXScript author
// BoolDef represent boolean object, contains only `true` or `false`

package ast

import (
	"strconv"
)

var _ Node = (*BoolDef)(nil)

// BoolDef represent the `true` or `false` value, store in the Value field
type BoolDef struct {
	Value bool
}

func (x *BoolDef) Any(handler RxeHandler) any {
	return x.Value
}

func (x *BoolDef) Type() CmdType {
	return BoolDefAstCmd
}

// Execute return the real value, such as true or false
// user should use `CheckSystemError` to check value or system error occurred.
func (x *BoolDef) Execute(e RxeHandler) any {
	return x
}

func (x *BoolDef) String() string {
	return "bool(" + strconv.FormatBool(x.Value) + ")"
}
