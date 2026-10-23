//
// The RXScript author
// MinusEqualDef represent -=

package ast

import (
	"math/big"
)

var _ Node = (*MinusEqualDef)(nil)

type MinusEqualDef struct {
	A, B Node
}

func (p *MinusEqualDef) Any(handler RxeHandler) any {
	return CheckSystemError(p.Execute(handler), handler).(Node).Any(handler)
}

func (p *MinusEqualDef) String() string {
	return "'-='"
}

func (p *MinusEqualDef) Execute(e RxeHandler) any {
	x1 := rawValue(p.A, e)
	x2 := rawValue(p.B, e)

	var ret Node

	switch x1 := x1.(type) {
	case int64:
		switch x2 := x2.(type) {
		case int64:
			ret = &IntDef{Value: x1 - x2}
		case float64:
			ret = &DoubleDef{Value: float64(x1) - x2}
		default:
			goto err
		}
	case float64:
		switch x2 := x2.(type) {
		case int64:
			ret = &DoubleDef{Value: x1 - float64(x2)}
		case float64:
			ret = &DoubleDef{Value: x1 - x2}
		default:
			goto err
		}
	case string:
		switch x2 := x2.(type) {
		case string:
			ab, ok := new(big.Int).SetString(x1, 10)
			if !ok {
				goto err
			}
			bb, ok := new(big.Int).SetString(x2, 10)
			if !ok {
				goto err
			}
			ret = &StringDef{Value: new(big.Int).Sub(ab, bb).String()}
		default:
			goto err
		}
	}

	switch stmt := p.A.(type) {
	case *GetAttributeDef:
		obj := CheckSystemError(stmt.Object.Execute(e), e)
		switch obj := obj.(type) {
		case *Class:
			obj.variables[stmt.Attribute] = ret
			return ret
		case *DictionaryDef:
			obj.Values[stmt.Attribute] = ret
			return ret
		default:
			return ThrowSystemErrorWithFormat("RXScript: cannot found object with name: '%s'", stmt.Object)
		}
	default:
		e.SetArgs(x1.(string), ret)
		return ret
	}

err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '-=' to types '%T' and '%T'", x1, x2)
}

func (p *MinusEqualDef) Type() CmdType {
	return MinusEqualDefAstCmd
}
