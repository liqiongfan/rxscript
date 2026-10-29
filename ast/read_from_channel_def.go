//
// The RXScript author
// ReadFromChannelDef represent <- xxx

package ast

var _ Node = (*ReadFromChannelDef)(nil)

type ReadFromChannelDef struct {
	Name Node
}

func (x *ReadFromChannelDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler)
}

func (x *ReadFromChannelDef) Type() CmdType {
	return ReadFromChannelDefAstCmd
}

func (x *ReadFromChannelDef) Execute(e RxeHandler) any {
	v := CheckSystemError(x.Name.Execute(e), e)
	switch d := v.(type) {
	case *ChannelDef:
		return <-d.Value
	case *TimeTicker:
		return <-d.Value.C
	case *TimeTimer:
		return <-d.Value.C
	default:
		return ThrowSystemError("RXScript: <- can only used on channel value")
	}
}

func (x *ReadFromChannelDef) String() string {
	return "object(<-channel)"
}
