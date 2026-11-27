//
// The RXScript author
// built-in functions

package ast

import (
	"bytes"
	"io"
)

type HttpRequestBody struct {
	FuncDef
}

func (p *HttpRequestBody) String() string {
	return "'httpRequestBody'"
}

func NewHttpRequestBody() *HttpRequestBody {
	return &HttpRequestBody{}
}

// Execute HttpRequestBody(httpResponse)
func (p *HttpRequestBody) Execute(rxe RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: httpRequestBody need 1 arguments")
	}

	switch v := p.Args[0].(type) {
	case *HttpRequest:
		body, _ := io.ReadAll(v.Value.Body)
		v.Value.Body = io.NopCloser(bytes.NewReader(body))
		return &StringDef{Value: string(body)}
	}

	return &NullDef{}
}
