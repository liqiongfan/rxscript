//
// The RXScript author
// BreakDef represent break, can only used in for statements

package ast

var _ Node = (*BreakDef)(nil)

type BreakDef struct {
}

func (x *BreakDef) Any(handler RxeHandler) any {
	return nil
}

func (x *BreakDef) Type() CmdType {
	return BreakDefAstCmd
}

func (x *BreakDef) Execute(e RxeHandler) any {

	return x
}

func (x *BreakDef) String() string {
	return "break"
}
