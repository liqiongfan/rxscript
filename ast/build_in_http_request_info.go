//
// The RXScript author
// built-in functions

package ast

type HttpRequestInfo struct {
	FuncDef
}

func (p *HttpRequestInfo) String() string {
	return "'httpRequestInfo'"
}

func NewHttpRequestInfo() *HttpRequestInfo {
	return &HttpRequestInfo{}
}

// Execute HttpRequestInfo(httpResponse)
func (p *HttpRequestInfo) Execute(rxe RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: httpRequestInfo need 1 arguments")
	}

	switch v := p.Args[0].(type) {
	case *HttpRequest:
		return &ArrayDef{Value: []Node{
			&StringDef{v.Value.Host},
			&StringDef{v.Value.Method},
			&StringDef{v.Value.Proto},
			// &StringDef{v.Value.Pattern},
			&StringDef{v.Value.RemoteAddr},
		}}
	}

	return &NullDef{}
}
