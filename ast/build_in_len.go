//
// The RXScript author
// built-in functions

package ast

type Len struct {
	FuncDef
}

func (p *Len) String() string {
	return "'len'"
}

func NewLen() *Len {
	return &Len{}
}

func (p *Len) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		arg := p.Args[0]
		switch v := arg.(type) {
		case *StringDef:
			return &IntDef{Value: int64(len(v.Value))}
		case *DictionaryDef:
			return &IntDef{Value: int64(len(v.Values))}
		case *ArrayDef:
			return &IntDef{Value: int64(len(v.Value))}
		}
	}
	return &IntDef{Value: 0}
}
