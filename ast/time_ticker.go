//
// The RXScript author
// ChannelDef represent channel object

package ast

import (
	"time"
)

var _ Node = (*TimeTicker)(nil)

type TimeTicker struct {
	Value *time.Ticker
}

func (x *TimeTicker) Any(handler RxeHandler) any {
	return x.Value
}

func (x *TimeTicker) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *TimeTicker) Execute(e RxeHandler) any {
	return x
}

func (x *TimeTicker) String() string {
	return "'time.Ticker'"
}
