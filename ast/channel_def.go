//
// The RXScript author
// ChannelDef represent channel object

package ast

var _ Node = (*ChannelDef)(nil)

type ChannelDef struct {
	Value chan any
}

func (x *ChannelDef) Any(handler RxeHandler) any {
	return x.Value
}

func (x *ChannelDef) Type() CmdType {
	return ChannelDefAstCmd
}

func (x *ChannelDef) Execute(e RxeHandler) any {
	return x
}

func (x *ChannelDef) String() string {
	return "'chan'"
}
