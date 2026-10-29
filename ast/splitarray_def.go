//
// The RXScript author
// SplitArrayDef represent []...

package ast

var _ Node = (*SplitArrayDef)(nil)

type SplitArrayDef struct {
	Value Node
}

func (x *SplitArrayDef) Any(handler RxeHandler) any {
	return (&ArrayDef{Value: CheckSystemError(x.Execute(handler), handler).([]Node)}).Any(handler)
}

func (x *SplitArrayDef) Type() CmdType {
	return SplitArrayDefAstCmd
}

func (x *SplitArrayDef) Execute(e RxeHandler) any {
	v := x.Value.Execute(e)
	switch v := v.(type) {
	case *ArrayDef:
		return v.Value
	default:
		return ThrowSystemErrorWithFormat("RXScript: ... can only be used on array. not '%s'", v)
	}
}

func (x *SplitArrayDef) String() string {
	return `object(...)`
}
