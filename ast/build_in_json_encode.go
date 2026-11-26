//
// The RXScript author
// built-in functions

package ast

type JsonEncode struct {
	FuncDef
}

func (p *JsonEncode) String() string {
	return "'jsonEncode'"
}

func NewJsonEncode() *JsonEncode {
	return &JsonEncode{}
}

func (p *JsonEncode) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: jsonEncode: need at least 1 argument for json encode.")
	}

	buff := RawAstNode(p.Args[0], rxe)

	switch v := buff.(type) {
	case Node:
		return jsonEncode(v, rxe)
	default:
		return ThrowSystemErrorWithFormat("RXScript: jsonEncode: decode data must be object, not '%T'", buff)
	}
}
