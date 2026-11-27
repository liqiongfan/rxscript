//
// The RXScript author
// built-in functions

package ast

import (
	"net/http"
)

type HttpServerListen struct {
	FuncDef
}

func (p *HttpServerListen) String() string {
	return "'httpServerListen'"
}

func NewHttpServerListen() *HttpServerListen {
	return &HttpServerListen{}
}

func (p *HttpServerListen) Execute(rxe RxeHandler) any {
	if len(p.Args) != 2 {
		return ThrowSystemError("RXScript: httpServerListen need at least 2 arguments.")
	}
	addr := p.Args[0]
	mux := p.Args[1]

	switch m := addr.(type) {
	case *StringDef:
		switch p := mux.(type) {
		case *HttpServerMux:
			return http.ListenAndServe(m.Value, p.Value)
		}
	}

	return ThrowSystemError("RXScript: unsupported mux handler")
}
