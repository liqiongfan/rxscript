//
// The RXScript author
// built-in functions

package ast

type HttpRequestPostForm struct {
	FuncDef
}

func (p *HttpRequestPostForm) String() string {
	return "'httpRequestPostForm'"
}

func NewHttpRequestPostForm() *HttpRequestPostForm {
	return &HttpRequestPostForm{}
}

// Execute HttpRequestPostForm(httpResponse)
func (p *HttpRequestPostForm) Execute(rxe RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: httpRequestPostForm need 1 arguments")
	}

	switch v := p.Args[0].(type) {
	case *HttpRequest:
		result := &DictionaryDef{Values: make(map[string]Node, len(v.Value.PostForm))}
		for key, values := range v.Value.PostForm {
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
