//
// The RXScript author
// ChannelDef represent channel object

package ast

import (
	"net/http"
)

var _ Node = (*HttpRequest)(nil)

type HttpRequest struct {
	Value *http.Request
}

func (x *HttpRequest) Any(handler RxeHandler) any {
	return x.Value
}

func (x *HttpRequest) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *HttpRequest) Execute(e RxeHandler) any {
	return x
}

func (x *HttpRequest) String() string {
	return "'http.request'"
}
