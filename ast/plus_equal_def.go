//
// The RXScript author
// PlusEqualDef represent +=

package ast

import (
	"math/big"
	"strconv"
)

var _ Node = (*PlusEqualDef)(nil)

type PlusEqualDef struct {
	A, B Node
}

func (p *PlusEqualDef) Any(handler RxeHandler) any {
	return CheckSystemError(p.Execute(handler), handler).(Node).Any(handler)
}

func (p *PlusEqualDef) String() string {
	return "'-='"
}

func (p *PlusEqualDef) Execute(e RxeHandler) any {
	x1 := rawValue(p.A, e)
	x2 := rawValue(p.B, e)

	var ret Node

	switch x1 := x1.(type) {
	case int64:
		switch x2 := x2.(type) {
		case int64:
			ret = &IntDef{Value: x1 + x2}
		case float64:
			ret = &DoubleDef{Value: float64(x1) + x2}
		case string:
			bb, ok := new(big.Rat).SetString(x2)
			if !ok {
				ret = &StringDef{Value: strconv.FormatInt(x1, 10) + x2}
				goto variableSet
			}
			v := new(big.Rat).Add(new(big.Rat).SetInt64(x1), bb)
			ret = &StringDef{Value: StripTrailingZeros(v.FloatString(FloatPrecisionWidth))}
		default:
			goto err
		}
	case float64:
		switch x2 := x2.(type) {
		case int64:
			ret = &DoubleDef{Value: x1 + float64(x2)}
		case float64:
			ret = &DoubleDef{Value: x1 + x2}
		case string:
			bb, ok := new(big.Rat).SetString(x2)
			if !ok {
				ret = &StringDef{Value: strconv.FormatFloat(x1, 'f', FloatPrecisionWidth, 64) + x2}
				goto variableSet
			}
			v := new(big.Rat).Add(new(big.Rat).SetFloat64(x1), bb)
			ret = &StringDef{Value: StripTrailingZeros(v.FloatString(FloatPrecisionWidth))}
		default:
			goto err
		}
	case string:
		switch x2 := x2.(type) {
		case string:
			ab, ok := new(big.Rat).SetString(x1)
			if !ok {
				ret = &StringDef{Value: x1 + x2}
				goto variableSet
			}
			bb, ok := new(big.Rat).SetString(x2)
			if !ok {
				ret = &StringDef{Value: x2 + x2}
				goto variableSet
			}
			v := new(big.Rat).Add(ab, bb)
			ret = &StringDef{Value: StripTrailingZeros(v.FloatString(FloatPrecisionWidth))}
		case int64:
			ab, ok := new(big.Rat).SetString(x1)
			if !ok {
				ret = &StringDef{Value: x1 + strconv.FormatInt(x2, 10)}
				goto variableSet
			}
			v := new(big.Rat).Add(new(big.Rat).SetInt64(x2), ab)
			ret = &StringDef{Value: StripTrailingZeros(v.FloatString(FloatPrecisionWidth))}
		case float64:
			ab, ok := new(big.Rat).SetString(x1)
			if !ok {
				ret = &StringDef{Value: x1 + strconv.FormatFloat(x2, 'f', FloatPrecisionWidth, 64)}
				goto variableSet
			}
			v := new(big.Rat).Add(new(big.Rat).SetFloat64(x2), ab)
			ret = &StringDef{Value: StripTrailingZeros(v.FloatString(FloatPrecisionWidth))}
		default:
			goto err
		}
	default:
		goto err
	}

variableSet:
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
	case *VariableDef:
		e.SetArgs(stmt.Name, ret)
		return ret
	default:
		name := rawValue(p.A, e)
		e.SetArgs(name.(string), ret)
		return ret
	}

err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '+=' to types '%T' and '%T'", x1, x2)
}

func (p *PlusEqualDef) Type() CmdType {
	return PlusEqualDefAstCmd
}
