//
// The RXScript author
// built-in functions

package ast

type HttpRequestURL struct {
	FuncDef
}

func (p *HttpRequestURL) String() string {
	return "'httpRequestURL'"
}

func NewHttpRequestURL() *HttpRequestURL {
	return &HttpRequestURL{}
}

// Execute HttpRequestURL(httpResponse)
func (p *HttpRequestURL) Execute(rxe RxeHandler) any {
	switch v := p.Args[0].(type) {
	case *HttpRequest:
		return &StringDef{v.Value.RequestURI}
	}
	return ThrowSystemError("RXScript: httpRequestURL need 1 arguments")
}
