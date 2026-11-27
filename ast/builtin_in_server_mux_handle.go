//
// The RXScript author
// built-in functions

package ast

import (
	"net/http"
)

type ServerMuxHandleFunc struct {
	FuncDef
}

func (p *ServerMuxHandleFunc) String() string {
	return "'httpServerMuxHandleFunc'"
}

func NewServerMuxHandleFunc() *ServerMuxHandleFunc {
	return &ServerMuxHandleFunc{}
}

func (p *ServerMuxHandleFunc) Execute(rxe RxeHandler) any {
	if len(p.Args) != 3 {
		return ThrowSystemError("RXScript: httpServerMuxHandleFunc need at least 3 arguments.")
	}
	mux := p.Args[0]
	patten := p.Args[1]
	handler := p.Args[2]

	switch m := mux.(type) {
	case *HttpServerMux:
		switch p := patten.(type) {
		case *StringDef:
			m.Value.HandleFunc(
				p.Value,
				func(w http.ResponseWriter, r *http.Request) {
					switch h := handler.(type) {
					case *ClassBodyDef:
						prevObj := rxe.GetCallingObject()
						defer rxe.SetCallingObject(prevObj)
						rxe.SetCallingObject(h.owner)
						h.Invoke([]any{
							&HttpResponseWriter{Value: w},
							&HttpRequest{Value: r},
						}, rxe)
					case *FunctionDef:
						h.Invoke([]any{
							&HttpResponseWriter{Value: w},
							&HttpRequest{Value: r},
						}, rxe)
					case *GoResource:
						(h.obj.(http.Handler)).ServeHTTP(w, r)
					}
				},
			)
			return &BoolDef{Value: true}
		}
	}

	return ThrowSystemError("RXScript: unsupported mux handler")
}
