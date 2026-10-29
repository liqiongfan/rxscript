//
// The RXScript author
// built-in functions

package ast

type ArrayShift struct {
	FuncDef
}

func (p *ArrayShift) String() string {
	return "'array_shift'"
}

func NewArrayShift() *ArrayShift {
	return &ArrayShift{}
}

func (p *ArrayShift) Execute(rxe RxeHandler) any {
	if len(p.Args) != 1 {
		return ThrowSystemError("RXScript: array_shift need 1 arguments")
	}

	arrayValue := p.Args[0]

	switch av := arrayValue.(type) {
	case *ArrayDef:
		length := len(av.Value)
		if length < 1 {
			return &NullDef{}
		}
		v := av.Value[0]
		av.Value = av.Value[1:length]
		return v
	default:
		return ThrowSystemErrorWithFormat("RXScript: array_shift first argument must be an array, not '%s'", av)
	}
}
