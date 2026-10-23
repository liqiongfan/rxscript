//
// The RXScript author
// NullDef represent null

package ast

var _ Node = (*NullDef)(nil)

type NullDef struct{}

func (x *NullDef) Any(handler RxeHandler) any {
	return nil
}

func (x *NullDef) Type() CmdType {
	return NullDefAstCmd
}

func (x *NullDef) Execute(e RxeHandler) any {
	return x
}

func (x *NullDef) String() string {
	return "'null'"
}
