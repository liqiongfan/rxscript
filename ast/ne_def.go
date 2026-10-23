//
// The RXScript author
// NeDef represent !=

package ast

var _ Node = (*NeDef)(nil)

type NeDef struct {
	A Node
	B Node
}

func (x *NeDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *NeDef) String() string {
	return "'!='"
}

func (x *NeDef) Type() CmdType {
	return NeDefAstCmd
}

func (x *NeDef) Execute(e RxeHandler) any {
	a := rawValue(x.A.Execute(e), e)
	b := rawValue(x.B.Execute(e), e)
	return &BoolDef{Value: a != b}
}
