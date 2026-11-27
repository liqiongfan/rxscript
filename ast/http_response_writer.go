//
// The RXScript author
// HttpResponseWriter represent Go's http.response writer

package ast

import (
	"net/http"
)

var _ Node = (*HttpResponseWriter)(nil)

type HttpResponseWriter struct {
	Value http.ResponseWriter
}

func (x *HttpResponseWriter) Any(handler RxeHandler) any {
	return x.Value
}

func (x *HttpResponseWriter) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *HttpResponseWriter) Execute(e RxeHandler) any {
	return x
}

func (x *HttpResponseWriter) String() string {
	return "'http(responseWriter)'"
}
