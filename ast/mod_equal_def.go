//
// The RXScript author
// ModEqualDef represent %=

package ast

import (
	"fmt"
	"math/big"
)

var _ Node = (*ModEqualDef)(nil)

type ModEqualDef struct {
	A Node
	B Node
}

func (x *ModEqualDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *ModEqualDef) String() string {
	return "'%='"
}

func (x *ModEqualDef) Type() CmdType {
	return ModEqualDefAstCmd
}

func (x *ModEqualDef) Execute(e RxeHandler) any {
	x1 := rawValue(x.A.Execute(e), e)
	x2 := rawValue(x.B.Execute(e), e)

	var ret Node

	switch a := x1.(type) {
	case int64:
		switch b := x2.(type) {
		case int64:
			ret = &IntDef{Value: a % b}
		case string:
			bb, ok := new(big.Int).SetString(b, 0)
			if !ok {
				goto err
			}
			ret = &StringDef{Value: StripTrailingZeros(new(big.Int).Mod(
				new(big.Int).SetInt64(a), bb).String())}
		}
	case string:
		switch b := x2.(type) {
		case int64:
			ab, ok := new(big.Int).SetString(a, 0)
			if !ok {
				goto err
			}
			ret = &StringDef{Value: StripTrailingZeros(new(big.Int).Mod(ab,
				new(big.Int).SetInt64(b)).String())}
		case string:
			ab, ok := new(big.Int).SetString(a, 0)
			if !ok {
				goto err
			}
			bb, ok := new(big.Int).SetString(b, 0)
			if !ok {
				goto err
			}
			ret = &StringDef{Value: StripTrailingZeros(new(big.Int).Mod(ab, bb).String())}
		}
	}

	switch stmt := x.A.(type) {
	case *GetAttributeDef:
		obj := CheckSystemError(stmt.Object.Execute(e), e)
		attr := CheckSystemError(rawValue(stmt.Attribute.Execute(e), e), e)
		attrName := fmt.Sprintf(`%v`, attr)
		switch obj := obj.(type) {
		case *Class:
			obj.variables[attrName] = ret
			return ret
		case *DictionaryDef:
			obj.Values[attrName] = ret
			return ret
		case *ArrayIndexDef:
			v := obj.Value.Execute(e)
			switch vv := v.(type) {
			case *ArrayDef:
				v := vv.Execute(e)
				if v == nil {
					return ThrowSystemError("RXScript: null value can't be indexed")
				}
				if obj.Index >= int64(len(v.(*ArrayDef).Value)) {
					return ThrowSystemErrorWithFormat("RXScript: index out of range, index: %d, total: %d", obj.Index, len(v.([]Node)))
				}
				v.(*ArrayDef).Value[obj.Index] = ret
				return ret
			default:
				return ThrowSystemErrorWithFormat("RXScript: value of '%v' can't be indexed", rawValue(vv, e))
			}
		default:
			return ThrowSystemErrorWithFormat("RXScript: cannot found object with name: '%s'", stmt.Object)
		}
	case *VariableDef:
		e.SetArgs(stmt.Name, ret)
		return ret
	default:
		name := rawValue(x.A, e)
		e.SetArgs(name.(string), ret)
		return ret
	}

err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '%%=' to types '%T' and '%T'", x1, x2)
}
