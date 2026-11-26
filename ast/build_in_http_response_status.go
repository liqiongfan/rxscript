//
// The RXScript author
// built-in functions

package ast

type HttpResponseStatus struct {
	FuncDef
}

func (p *HttpResponseStatus) String() string {
	return "'httpResponseStatus'"
}

func NewHttpResponseStatus() *HttpResponseStatus {
	return &HttpResponseStatus{}
}

// Execute httpResponseStatus(httpResponse)
func (p *HttpResponseStatus) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		switch hr := p.Args[0].(type) {
		case *HttpResponseRequest:
			return &IntDef{Value: int64(hr.Value.StatusCode)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: method need httpResponse, not '%s'", hr)
		}
	}
	return ThrowSystemError("RXScript: httpResponseStatus need 1 arguments")
}
