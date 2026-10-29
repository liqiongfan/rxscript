//
// The RXScript author
// built-in functions

package ast

import (
	"math/big"
)

type ToInt struct {
	FuncDef
}

func (p *ToInt) String() string {
	return "'toInt'"
}

func NewToInt() *ToInt {
	return &ToInt{}
}

func (p *ToInt) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: toInt: need at least 1 argument.")
	}

	v := p.Args[0]
	switch vv := v.(type) {
	case *DoubleDef:
		return &IntDef{Value: int64(vv.Value)}
	case *IntDef:
		return &IntDef{Value: vv.Value}
	case *StringDef:
		val, ok := new(big.Int).SetString(vv.Value, 0)
		if !ok || (!val.IsInt64() && !val.IsUint64()) {
			goto err
		}
		return &IntDef{Value: val.Int64()}
	case *BoolDef:
		if vv.Value {
			return &IntDef{Value: 1}
		}
		return &IntDef{Value: 0}
	}

err:
	return ThrowSystemErrorWithFormat("RXScript: toInt: invalid argument, expect number not '%s'", v)
}
