//
// The RXScript author
// ArrayDef represent array/list object

package ast

import (
	"fmt"
)

var _ Node = (*ArrayDef)(nil)

// ArrayDef array value
// Each value of the array list must be ast.AstNode
// if you want to get the raw value, use ast.RawAstNode method
// array value is: [ 1, 20.0, null, false, "hello" ]
type ArrayDef struct {
	Value []Node
}

func (x *ArrayDef) Any(handler RxeHandler) any {
	ret := make([]any, 0, len(x.Value))
	for _, v := range x.Value {
		ret = append(ret, v.Any(handler))
	}
	return ret
}

// Type ArrayDefAstCmd
func (x *ArrayDef) Type() CmdType {
	return ArrayDefAstCmd
}

// Execute return the raw value of the array list
// user should use `CheckSystemError` to check value or system error occurred.
func (x *ArrayDef) Execute(e RxeHandler) any {
	return x
}

func (x *ArrayDef) String() string {
	v := "array["
	for i, n := range x.Value {
		v += fmt.Sprintf("%s", n)
		if i != len(x.Value)-1 {
			v += ","
		}
	}
	return v + "]"
}
