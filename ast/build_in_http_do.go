//
// The RXScript author
// built-in functions

package ast

import (
	"io"
	"net/http"
)

type HttpDoFunc struct {
	FuncDef
}

func (p *HttpDoFunc) String() string {
	return "'httpDo'"
}

func NewHttpDoFunc() *HttpDoFunc {
	return &HttpDoFunc{}
}

// Execute httpDo(httpRequest)
func (p *HttpDoFunc) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		switch hr := p.Args[0].(type) {
		case *HttpRequest:
			reply, err := http.DefaultClient.Do(hr.Value)
			if err != nil {
				return ThrowSystemError(err)
			}
			defer reply.Body.Close()
			response, err := io.ReadAll(reply.Body)
			if err != nil {
				return ThrowSystemError(err)
			}
			return &StringDef{Value: string(response)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: method need httpRequest, not '%s'", hr)
		}
	}
	return ThrowSystemError("RXScript: httpDo need 1 arguments")
}
