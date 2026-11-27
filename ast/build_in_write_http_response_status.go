//
// The RXScript author
// built-in functions

package ast

type WriteHttpResponseCode struct {
	FuncDef
}

func (p *WriteHttpResponseCode) String() string {
	return "'writeHttpResponseCode'"
}

func NewWriteHttpResponseCode() *WriteHttpResponseCode {
	return &WriteHttpResponseCode{}
}

func (p *WriteHttpResponseCode) Execute(rxe RxeHandler) any {
	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: writeHttpResponseCode need ResponseWriter & statusCode to write.")
	}
	writer := p.Args[0]
	code := p.Args[1]

	var statusCode int64

	switch l := code.(type) {
	case *IntDef:
		statusCode = l.Value
	default:
		statusCode = 200 // default value is OK
	}

	switch d := writer.(type) {
	case *HttpResponseWriter:
		d.Value.WriteHeader(int(statusCode))
		return &BoolDef{Value: true}
	default:
		return ThrowSystemErrorWithFormat("RXScript: writeHttpResponseCode need ResponseWriter object to write statusCode.")
	}
}
