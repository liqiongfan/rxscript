//
// The RXScript author
// EqualDoubleDef represent += in math.

package ast

var _ Node = (*EqualDoubleDef)(nil)

type EqualDoubleDef struct {
	A Node
	B Node
}

func (x *EqualDoubleDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *EqualDoubleDef) String() string {
	return "'=='"
}

func (x *EqualDoubleDef) Type() CmdType {
	return EqualDoubleDefAstCmd
}

func (x *EqualDoubleDef) Execute(e RxeHandler) any {
	a := rawValue(x.A.Execute(e), e)
	b := rawValue(x.B.Execute(e), e)
	return &BoolDef{Value: a == b}
}
