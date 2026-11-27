//
// The RXScript author
// built-in functions

package ast

type HttpRequestHeader struct {
	FuncDef
}

func (p *HttpRequestHeader) String() string {
	return "'httpRequestHeader'"
}

func NewHttpRequestHeader() *HttpRequestHeader {
	return &HttpRequestHeader{}
}

// Execute HttpRequestHeader(httpResponse)
func (p *HttpRequestHeader) Execute(rxe RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: httpRequestHeader need 1 arguments")
	}

	switch v := p.Args[0].(type) {
	case *HttpRequest:
		dict := &DictionaryDef{Values: make(map[string]Node, 32)}
		for key, name := range v.Value.Header {
			values := &ArrayDef{Value: make([]Node, len(name))}
			for i, value := range name {
				values.Value[i] = &StringDef{Value: value}
			}
			dict.Values[key] = values
		}
		return dict
	}

	return &NullDef{}
}
