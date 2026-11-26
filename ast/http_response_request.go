//
// The RXScript author
// HttpRequest represent Go's http.Request

package ast

import (
	"net/http"
)

var _ Node = (*HttpResponseRequest)(nil)

type HttpResponseRequest struct {
	Value *http.Response
}

func (x *HttpResponseRequest) Any(handler RxeHandler) any {
	return x.Value
}

func (x *HttpResponseRequest) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *HttpResponseRequest) Execute(e RxeHandler) any {
	return x
}

func (x *HttpResponseRequest) String() string {
	return "'http(response)'"
}
