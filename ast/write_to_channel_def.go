//
// The RXScript author
// WriteToChannelDef represent c <-

package ast

var _ Node = (*WriteToChannelDef)(nil)

type WriteToChannelDef struct {
	Name  Node
	Value Node
}

func (x *WriteToChannelDef) Any(handler RxeHandler) any {
	return nil
}

func (x *WriteToChannelDef) Type() CmdType {
	return WriteToChannelDefAstCmd
}

func (x *WriteToChannelDef) Execute(e RxeHandler) any {
	v := CheckSystemError(x.Name.Execute(e), e)
	switch d := v.(type) {
	case *ChannelDef:
		d.Value <- CheckSystemError(x.Value.Execute(e), e)
		return d.Value
	default:
		return ThrowSystemError("RXScript: <- can only used on channel value")
	}
}

func (x *WriteToChannelDef) String() string {
	return "object(channel<-)"
}
