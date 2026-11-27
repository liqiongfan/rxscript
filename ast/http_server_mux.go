//
// The RXScript author
// HttpServerMux represent Go's http.server mux

package ast

import (
	"net/http"
)

var _ Node = (*HttpServerMux)(nil)

type HttpServerMux struct {
	Value *http.ServeMux
}

func (x *HttpServerMux) Any(handler RxeHandler) any {
	return x.Value
}

func (x *HttpServerMux) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *HttpServerMux) Execute(e RxeHandler) any {
	return x
}

func (x *HttpServerMux) String() string {
	return "'http(serverMux)'"
}
