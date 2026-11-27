//
// The RXScript author
// built-in functions

package ast

import (
	"net/http"
)

type ServerMuxHandler struct {
	FuncDef
}

func (p *ServerMuxHandler) String() string {
	return "'httpServerMuxHandler'"
}

func NewServerMuxHandler() *ServerMuxHandler {
	return &ServerMuxHandler{}
}

func (p *ServerMuxHandler) Execute(rxe RxeHandler) any {
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
			switch h := handler.(type) {
			case *GoResource:
				m.Value.Handle(p.Value, h.obj.(http.Handler))
			}
			return &BoolDef{Value: true}
		}
	}

	return ThrowSystemError("RXScript: unsupported mux handler")
}
