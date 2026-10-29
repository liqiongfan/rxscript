//
// The RXScript author
// ReturnDef represent return v

package ast

var _ Node = (*ReturnDef)(nil)

type ReturnDef struct {
	Value any
}

func (x *ReturnDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *ReturnDef) String() string {
	return "object(return)"
}

func (x *ReturnDef) Type() CmdType {
	return ReturnDefAstCmd
}

func (x *ReturnDef) Execute(e RxeHandler) any {
	switch xv := x.Value.(type) {
	case *IntDef, *StringDef, *BoolDef, *NullDef, *DoubleDef:
		return x
	case Node:
		return xv.Execute(e)
	default:
		return ThrowSystemErrorWithFormat("RXScript: unsupported type: '%T'", xv)
	}
}
