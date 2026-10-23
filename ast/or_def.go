//
// The RXScript author
// OrDef represent ||

package ast

var _ Node = (*OrDef)(nil)

type OrDef struct {
	A Node
	B Node
}

func (x *OrDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *OrDef) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (x *OrDef) String() string {
	return "'||'"
}

func (x *OrDef) Type() CmdType {
	return OrDefAstCmd
}

func (x *OrDef) Execute(e RxeHandler) any {
	a := convertToBool(x.A.Execute(e), e)
	b := convertToBool(x.B.Execute(e), e)
	return &BoolDef{Value: a || b}
}
