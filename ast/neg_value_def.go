//
// The RXScript author
// NegValueDef represent -10

package ast

import "math/big"

var _ Node = (*NegValueDef)(nil)

type NegValueDef struct {
	Value Node
}

func (x *NegValueDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *NegValueDef) String() string {
	return "'-" + x.Value.String() + "'"
}

func (x *NegValueDef) Type() CmdType {
	return NegValueDefAstCmd
}

func (x *NegValueDef) Execute(e RxeHandler) any {
	v := rawValue(x.Value.Execute(e), e)
	switch v := v.(type) {
	case int64:
		return &IntDef{Value: v * -1}
	case float64:
		return &DoubleDef{Value: v * -1}
	case string:
		ab, ok := new(big.Rat).SetString(v)
		if !ok {
			goto err
		}
		r := new(big.Rat).Mul(new(big.Rat).SetInt64(-1), ab)
		if r.IsInt() {
			return &IntDef{Value: r.Num().Int64()}
		}
		return &StringDef{Value: StripTrailingZeros(r.FloatString(FloatPrecisionWidth))}
	}
err:
	return ThrowSystemErrorWithFormat("RXScript: unknown negative type: '%s'", v)
}
