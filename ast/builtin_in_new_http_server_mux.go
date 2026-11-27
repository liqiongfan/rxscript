//
// The RXScript author
// built-in functions

package ast

import (
	"net/http"
)

type NewHttpServerMux struct {
	FuncDef
}

func (p *NewHttpServerMux) String() string {
	return "'newHttpServerMux'"
}

func NewNewHttpServerMux() *NewHttpServerMux {
	return &NewHttpServerMux{}
}

func (p *NewHttpServerMux) Execute(rxe RxeHandler) any {
	return &HttpServerMux{Value: http.NewServeMux()}
}
