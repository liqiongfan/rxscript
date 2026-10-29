//
// The RXScript author
// built-in functions

package ast

type ArrayPop struct {
	FuncDef
}

func (p *ArrayPop) String() string {
	return "'arrayPop'"
}

func NewArrayPop() *ArrayPop {
	return &ArrayPop{}
}

func (p *ArrayPop) Execute(rxe RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: array_pop need 1 arguments")
	}

	arrayValue := p.Args[0]

	switch av := arrayValue.(type) {
	case *ArrayDef:
		length := len(av.Value)
		if length < 1 {
			return &NullDef{}
		}
		v := av.Value[length-1]
		av.Value = av.Value[0 : length-1]
		return v
	default:
		return ThrowSystemErrorWithFormat("RXScript: array_pop first argument must be an array, not '%s'", av)
	}
}
