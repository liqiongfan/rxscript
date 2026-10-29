//
// The RXScript author
// PlusDoubleDef represent ++

package ast

import "math/big"

var _ Node = (*PlusDoubleDef)(nil)

type PlusDoubleDef struct {
	A Node
}

func (p *PlusDoubleDef) Any(handler RxeHandler) any {
	return CheckSystemError(p.Execute(handler), handler).(Node).Any(handler)
}

func (p *PlusDoubleDef) String() string {
	return "'++'"
}

func (p *PlusDoubleDef) Execute(e RxeHandler) any {
	var ret, ov Node
	x := rawValue(p.A, e)

	switch n := x.(type) {
	case int64:
		ov = &IntDef{Value: n}
		ret = &IntDef{Value: n + 1}
	case float64:
		ov = &DoubleDef{Value: n}
		ret = &DoubleDef{Value: n + 1}
	case string:
		na, ok := new(big.Rat).SetString(n)
		if !ok {
			return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '++' to types '%T'", x)
		}
		r := new(big.Rat).Add(na, new(big.Rat).SetInt64(1))
		ret = &StringDef{Value: StripTrailingZeros(r.FloatString(FloatPrecisionWidth))}
		ov = &StringDef{Value: n}
	default:
		return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '++' to types '%T'", x)
	}

	switch stmt := p.A.(type) {
	case *GetAttributeDef:
		obj := CheckSystemError(stmt.Object.Execute(e), e)
		switch obj := obj.(type) {
		case *Class:
			obj.variables[stmt.Attribute] = ret
			return ov
		case *DictionaryDef:
			obj.Values[stmt.Attribute] = ret
			return ov
		default:
			return ThrowSystemErrorWithFormat("RXScript: cannot found object with name: '%s'", stmt.Object)
		}
	case *VariableDef:
		e.SetArgs(stmt.Name, ret)
		return ov
	default:
		name := rawValue(p.A, e)
		e.SetArgs(name.(string), ret)
		return ov
	}
}

func (p *PlusDoubleDef) Type() CmdType {
	return PlusDoubleDefAstCmd
}
