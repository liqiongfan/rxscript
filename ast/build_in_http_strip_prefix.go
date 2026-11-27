//
// The RXScript author
// built-in functions

package ast

import "net/http"

type HttpStripPrefix struct {
	FuncDef
}

func (p *HttpStripPrefix) String() string {
	return "'httpStripPrefix'"
}

func NewHttpStripPrefix() *HttpStripPrefix {
	return &HttpStripPrefix{}
}

// Execute HttpStripPrefix(httpResponse)
func (p *HttpStripPrefix) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 1 {
		return ThrowSystemError("RXScript: httpStripPrefix need 2 arguments")
	}

	prefix := p.Args[0]
	fs := p.Args[1]

	switch hr := prefix.(type) {
	case *StringDef:
		switch f := fs.(type) {
		case *GoResource:
			return &GoResource{obj: http.StripPrefix(hr.Value, f.obj.(http.Handler))}
		}
	}
	return ThrowSystemErrorWithFormat("RXScript: httpStripPrefix need string, not '%T'", prefix)
}
