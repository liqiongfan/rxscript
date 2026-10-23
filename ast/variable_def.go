//
// The RXScript author
// VariableDef represent b

package ast

var _ Node = (*VariableDef)(nil)

type VariableDef struct {
	Name string
}

func (x *VariableDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *VariableDef) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (x *VariableDef) String() string {
	return x.Name
}

func (x *VariableDef) Execute(e RxeHandler) any {
	return fetchFromRxe(e, x.Name)
}

func (x *VariableDef) Type() CmdType {
	return VariableDefAstCmd
}
