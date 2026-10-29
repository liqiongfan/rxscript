//
// The RXScript author
// DivideDef represent /

package ast

import (
	"math/big"
)

var _ Node = (*DivideDef)(nil)

type DivideDef struct {
	A Node
	B Node
}

func (x *DivideDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *DivideDef) String() string {
	return "'" + x.A.String() + "/" + x.B.String() + "'"
}

func (x *DivideDef) Type() CmdType {
	return DivideDefAstCmd
}

func (x *DivideDef) Execute(e RxeHandler) any {
	x1 := rawValue(x.A.Execute(e), e)
	x2 := rawValue(x.B.Execute(e), e)

	switch a := x1.(type) {
	case int64:
		switch b := x2.(type) {
		case int64:
			return &IntDef{Value: a / b}
		case float64:
			return &DoubleDef{Value: float64(a) / b}
		case string:
			bb, ok := new(big.Rat).SetString(b)
			if !ok {
				goto err
			}
			return &StringDef{Value: StripTrailingZeros(new(big.Rat).Quo(
				new(big.Rat).SetInt64(a), bb).FloatString(FloatPrecisionWidth))}
		}
	case float64:
		switch b := x2.(type) {
		case int64:
			return &DoubleDef{Value: a / float64(b)}
		case float64:
			return &DoubleDef{Value: a / b}
		case string:
			bb, ok := new(big.Rat).SetString(b)
			if !ok {
				goto err
			}
			return &StringDef{Value: StripTrailingZeros(new(big.Rat).Quo(
				new(big.Rat).SetFloat64(a), bb).FloatString(FloatPrecisionWidth))}
		}
	case string:
		switch b := x2.(type) {
		case int64:
			ab, ok := new(big.Rat).SetString(a)
			if !ok {
				goto err
			}
			return &StringDef{Value: StripTrailingZeros(new(big.Rat).Quo(ab,
				new(big.Rat).SetInt64(b)).FloatString(FloatPrecisionWidth))}
		case float64:
			ab, ok := new(big.Rat).SetString(a)
			if !ok {
				goto err
			}
			return &StringDef{Value: StripTrailingZeros(new(big.Rat).Quo(ab,
				new(big.Rat).SetFloat64(b)).FloatString(FloatPrecisionWidth))}
		case string:
			ab, ok := new(big.Rat).SetString(a)
			if !ok {
				goto err
			}
			bb, ok := new(big.Rat).SetString(b)
			if !ok {
				goto err
			}
			return &StringDef{Value: StripTrailingZeros(new(big.Rat).Quo(ab, bb).FloatString(FloatPrecisionWidth))}
		}
	}

err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '/' to types '%T' and '%T'", x1, x2)
}
