//
// The RXScript author
// built-in functions

package ast

import (
	"net/http"
)

type HttpServerTLSListen struct {
	FuncDef
}

func (p *HttpServerTLSListen) String() string {
	return "'httpServerTLSListen'"
}

func NewHttpServerTLSListen() *HttpServerTLSListen {
	return &HttpServerTLSListen{}
}

func (p *HttpServerTLSListen) Execute(rxe RxeHandler) any {
	if len(p.Args) != 2 {
		return ThrowSystemError("RXScript: httpServerTLSListen need at least 4 arguments.")
	}
	addr := p.Args[0]
	mux := p.Args[1]
	cert := p.Args[2]
	key := p.Args[3]

	switch cert.(type) {
	case *StringDef:
	default:
		return ThrowSystemError("RXScript: httpServerTLSListen need a string cert file path.")
	}

	switch key.(type) {
	case *StringDef:
	default:
		return ThrowSystemError("RXScript: httpServerTLSListen need a string key file path.")
	}

	switch m := addr.(type) {
	case *StringDef:
		switch p := mux.(type) {
		case *HttpServerMux:
			return http.ListenAndServeTLS(m.Value, cert.(*StringDef).Value, key.(*StringDef).Value, p.Value)
		}
	}

	return ThrowSystemError("RXScript: unsupported mux handler")
}
