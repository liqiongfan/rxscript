//
// The RXScript author
// built-in functions

package ast

import "fmt"

type SetHttpHeaderFunc struct {
	FuncDef
}

func (p *SetHttpHeaderFunc) String() string {
	return "'setHttpHeader'"
}

func NewSetHttpHeaderFunc() *SetHttpHeaderFunc {
	return &SetHttpHeaderFunc{}
}

// Execute setHttpHeader(httpRequest, headerName, headerValue)
func (p *SetHttpHeaderFunc) Execute(rxe RxeHandler) any {
	if len(p.Args) > 2 {
		switch hr := p.Args[0].(type) {
		case *HttpRequest:
			switch headerName := p.Args[1].(type) {
			case *StringDef:
				switch headerValue := p.Args[2].(type) {
				case *StringDef:
					hr.Value.Header.Add(headerName.Value, headerValue.Value)
					return &BoolDef{Value: true}
				case *IntDef:
					hr.Value.Header.Add(headerName.Value, fmt.Sprintf("%d", headerValue.Value))
					return &BoolDef{Value: true}
				case *DoubleDef:
					hr.Value.Header.Add(headerName.Value, fmt.Sprintf("%f", headerValue.Value))
					return &BoolDef{Value: true}
				default:
					return ThrowSystemErrorWithFormat("RXScript: Invalid header value, '%s'", headerValue)
				}
			default:
				return ThrowSystemErrorWithFormat("RXScript: method need string, not '%s'", headerName)
			}
		default:
			return ThrowSystemErrorWithFormat("RXScript: method need httpRequest, not '%s'", hr)
		}
	}
	return ThrowSystemError("RXScript: setHttpHeader need at least 3 arguments")
}
