//
// The RXScript author
// built-in functions

package ast

type HttpRequestForm struct {
	FuncDef
}

func (p *HttpRequestForm) String() string {
	return "'httpRequestForm'"
}

func NewHttpRequestForm() *HttpRequestForm {
	return &HttpRequestForm{}
}

// Execute HttpRequestForm(httpResponse)
func (p *HttpRequestForm) Execute(rxe RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: httpRequestForm need 1 arguments")
	}

	switch v := p.Args[0].(type) {
	case *HttpRequest:
		result := &DictionaryDef{Values: make(map[string]Node, len(v.Value.Form))}
		for key, values := range v.Value.Form {
			val := &ArrayDef{Value: make([]Node, 0, len(values))}
			for _, value := range values {
				val.Value = append(val.Value, &StringDef{Value: value})
			}
			result.Values[key] = val
		}
		return result
	}

	return &NullDef{}
}
