//
// The RXScript author
// built-in functions

package ast

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpRequestFunc struct {
	FuncDef
}

func (p *HttpRequestFunc) String() string {
	return "'HttpRequest'"
}

func NewHttpRequestFunc() *HttpRequestFunc {
	return &HttpRequestFunc{}
}

// Execute newRequest(method, url, data)
func (p *HttpRequestFunc) Execute(rxe RxeHandler) any {
	if len(p.Args) > 1 {
		switch method := p.Args[0].(type) {
		case *StringDef:
			switch url := p.Args[1].(type) {
			case *StringDef:
				var buff []byte
				var err error
				if len(p.Args) > 2 {
					switch data := p.Args[2].(type) {
					case *NullDef:
						break
					case Node:
						buff, err = json.Marshal(data.Any(rxe))
						if err != nil {
							return ThrowSystemError(err)
						}
					default:
						panic("unreachable")
					}
				}
				if len(buff) > 0 {
					source, err := http.NewRequest(method.Value, url.Value, bytes.NewBuffer(buff))
					if err != nil {
						return ThrowSystemError(err)
					}
					return &HttpRequest{Value: source}
				}
				source, err := http.NewRequest(method.Value, url.Value, nil)
				if err != nil {
					return ThrowSystemError(err)
				}
				return &HttpRequest{Value: source}
			default:
				return ThrowSystemErrorWithFormat("RXScript: method need string, not '%s'", url)
			}
		default:
			return ThrowSystemErrorWithFormat("RXScript: method need string, not '%s'", method)
		}
	}

	return ThrowSystemError("RXScript: newHttp need at least two arguments")
}
