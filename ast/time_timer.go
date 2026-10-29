//
// The RXScript author
// ChannelDef represent channel object

package ast

import (
	"time"
)

var _ Node = (*TimeTimer)(nil)

type TimeTimer struct {
	Value *time.Timer
}

func (x *TimeTimer) Any(handler RxeHandler) any {
	return x.Value
}

func (x *TimeTimer) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *TimeTimer) Execute(e RxeHandler) any {
	return x
}

func (x *TimeTimer) String() string {
	return "object(timer)"
}
