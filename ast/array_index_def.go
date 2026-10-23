//
// The RXScript author
// ArrayIndexDef represent list[x] object

package ast

import "fmt"

var _ Node = (*ArrayIndexDef)(nil)

// ArrayIndexDef array value index operator
// Value can be any value, even ArrayIndexDef
// Index is the index number
// such as: v = [1,2,3][1]
// Value is [1,2,3]
// Index is 1
// or, v = [[1, 2],[2, 3]][1][1]
// Value is [[1, 2],[2, 3]][1]
// Index is 1
type ArrayIndexDef struct {
	Value Node
	Index int64
}

func (x *ArrayIndexDef) Any(handler RxeHandler) any {
	nodes := x.Value.Any(handler).([]Node)
	if x.Index >= int64(len(nodes)) {
		return ThrowSystemErrorWithFormat("RXScript: index out of range, index: %d, total: %d", x.Index, len(nodes))
	}
	return nodes[x.Index].Any(handler)
}

func (x *ArrayIndexDef) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (x *ArrayIndexDef) String() string {
	return fmt.Sprintf(`'%s[%d]'`, x.Value.String(), x.Index)
}

// Type ast.ArrayIndexDefAstCmd
func (x *ArrayIndexDef) Type() CmdType {
	return ArrayIndexDefAstCmd
}

// Execute the ArrayIndexDef and return the value of the value
// return value must be ast.AstNode
// user should use `CheckSystemError` to check value or system error occurred.
func (x *ArrayIndexDef) Execute(e RxeHandler) any {
	v := x.Value.Execute(e)
	switch vv := v.(type) {
	case *StringDef:
		if x.Index >= int64(len(vv.Value)) {
			return ThrowSystemErrorWithFormat("RXScript: index out of range, index: %d, total: %d", x.Index, len(vv.Value))
		}
		return &StringDef{Value: string(vv.Value[x.Index])}
	case *ArrayDef:
		v := vv.Execute(e)
		if v == nil {
			return ThrowSystemError("RXScript: null value can't be indexed")
		}
		if x.Index >= int64(len(v.(*ArrayDef).Value)) {
			return ThrowSystemErrorWithFormat("RXScript: index out of range, index: %d, total: %d", x.Index, len(v.([]Node)))
		}
		return v.(*ArrayDef).Value[x.Index]
	default:
		return ThrowSystemErrorWithFormat("RXScript: value of '%v' can't be indexed", rawValue(vv, e))
	}
}
