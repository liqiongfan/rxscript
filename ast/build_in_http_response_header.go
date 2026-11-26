//
// The RXScript author
// built-in functions

package ast

type HttpResponseHeader struct {
	FuncDef
}

func (p *HttpResponseHeader) String() string {
	return "'httpResponseHeader'"
}

func NewHttpResponseHeader() *HttpResponseHeader {
	return &HttpResponseHeader{}
}

// Execute HttpResponseHeader(httpResponse)
func (p *HttpResponseHeader) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		switch hr := p.Args[0].(type) {
		case *HttpResponseRequest:
			result := &DictionaryDef{Values: make(map[string]Node, len(hr.Value.Header))}
			for key, values := range hr.Value.Header {
				val := &ArrayDef{Value: make([]Node, 0, len(values))}
				for _, value := range values {
					val.Value = append(val.Value, &StringDef{Value: value})
				}
				result.Values[key] = val
			}
			return result
		default:
			return ThrowSystemErrorWithFormat("RXScript: method need httpResponse, not '%s'", hr)
		}
	}
	return ThrowSystemError("RXScript: httpResponseHeader need 1 arguments")
}
