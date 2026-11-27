//
// The RXScript author
// built-in functions

package ast

type HttpRequestPathValue struct {
	FuncDef
}

func (p *HttpRequestPathValue) String() string {
	return "'httpRequestPathValue'"
}

func NewHttpRequestPathValue() *HttpRequestPathValue {
	return &HttpRequestPathValue{}
}

// Execute HttpRequestPathValue(httpResponse)
func (p *HttpRequestPathValue) Execute(rxe RxeHandler) any {
	if len(p.Args) != 2 {
		return ThrowSystemError("RXScript: httpRequestPathValue need 2 arguments")
	}

	switch v := p.Args[0].(type) {
	case *HttpRequest:
		switch name := p.Args[1].(type) {
		case *StringDef:
			return &StringDef{v.Value.PathValue(name.Value)}
		}
	}

	return &NullDef{}
}
