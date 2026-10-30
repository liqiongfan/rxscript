//
// The RXScript author
// NotDef represent !

package ast

var _ Node = (*NotDef)(nil)

type NotDef struct {
	A Node
}

func (x *NotDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *NotDef) String() string {
	return "'!'"
}

func (x *NotDef) Type() CmdType {
	return NotDefAstCmd
}

func (x *NotDef) Execute(e RxeHandler) any {
	v := convertToBool(x.A, e)
	return &BoolDef{Value: !v}
}
