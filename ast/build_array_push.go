//
// The RXScript author
// built-in functions

package ast

type ArrayPush struct {
	FuncDef
}

func (p *ArrayPush) String() string {
	return "'arrayPush'"
}

func NewArrayPush() *ArrayPush {
	return &ArrayPush{}
}

func (p *ArrayPush) Execute(rxe RxeHandler) any {
	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: arrayPush need at least 2 arguments")
	}

	arrayValue := p.Args[0]
	values := p.Args[1:]

	switch av := arrayValue.(type) {
	case *ArrayDef:
		for _, value := range values {
			av.Value = append(av.Value, value.(Node))
		}
	default:
		return ThrowSystemErrorWithFormat("RXScript: arrayPush first argument must be an array, not '%s'", av)
	}

	return &NullDef{}
}
