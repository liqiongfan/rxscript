//
// The RXScript author
// built-in functions

package ast

import "net/http"

type HttpFileServer struct {
	FuncDef
}

func (p *HttpFileServer) String() string {
	return "'httpFileServer'"
}

func NewHttpFileServer() *HttpFileServer {
	return &HttpFileServer{}
}

// Execute HttpFileServer(httpResponse)
func (p *HttpFileServer) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: HttpFileServer need 1 arguments")
	}

	switch hr := p.Args[0].(type) {
	case *StringDef:
		return &GoResource{obj: http.FileServer(http.Dir(hr.Value))}
	default:
		return ThrowSystemErrorWithFormat("RXScript: httpFileServer need string, not '%s'", hr)
	}
}
