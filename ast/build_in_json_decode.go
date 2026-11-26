//
// The RXScript author
// built-in functions

package ast

type JsonDecode struct {
	FuncDef
}

func (p *JsonDecode) String() string {
	return "'jsonDecode'"
}

func NewJsonDecode() *JsonDecode {
	return &JsonDecode{}
}

func (p *JsonDecode) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: jsonDecode: need at least 1 argument for json decode.")
	}

	buff := p.Args[0]

	switch v := buff.(type) {
	case *StringDef:
		return jsonDecode(v.Value)
	case string:
		return jsonDecode(v)
	default:
		return ThrowSystemErrorWithFormat("RXScript: jsonDecode: decode data must be string, not '%T'", buff)
	}
}
