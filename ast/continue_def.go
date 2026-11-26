//
// The RXScript author
// ContinueDef represent break, can only used in for statements

package ast

var _ Node = (*ContinueDef)(nil)

type ContinueDef struct {
}

func (x *ContinueDef) Any(handler RxeHandler) any {
	return `{}`
}

func (x *ContinueDef) Type() CmdType {
	return ContinueDefAstCmd
}

func (x *ContinueDef) Execute(e RxeHandler) any {

	return x
}

func (x *ContinueDef) String() string {
	return "continue"
}
