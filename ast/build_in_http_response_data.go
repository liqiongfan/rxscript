//
// The RXScript author
// built-in functions

package ast

import "io"

type HttpResponseData struct {
	FuncDef
}

func (p *HttpResponseData) String() string {
	return "'httpResponseData'"
}

func NewHttpResponseData() *HttpResponseData {
	return &HttpResponseData{}
}

// Execute HttpResponseData(httpResponse)
func (p *HttpResponseData) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		switch hr := p.Args[0].(type) {
		case *HttpResponseRequest:
			defer hr.Value.Body.Close()
			buf, err := io.ReadAll(hr.Value.Body)
			if err != nil {
				return ThrowSystemError(err)
			}
			return &StringDef{Value: string(buf)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: method need httpResponse, not '%s'", hr)
		}
	}
	return ThrowSystemError("RXScript: httpResponseData need 1 arguments")
}
