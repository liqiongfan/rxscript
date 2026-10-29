//
// The RXScript author
// DivideDef represent /

package ast

import (
	"math/big"
)

var _ Node = (*ModDef)(nil)

type ModDef struct {
	A Node
	B Node
}

func (x *ModDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *ModDef) String() string {
	return "'%'"
}

func (x *ModDef) Type() CmdType {
	return ModDefAstCmd
}

func (x *ModDef) Execute(e RxeHandler) any {
	x1 := rawValue(x.A.Execute(e), e)
	x2 := rawValue(x.B.Execute(e), e)

	switch a := x1.(type) {
	case int64:
		switch b := x2.(type) {
		case int64:
			return &IntDef{Value: a % b}
		case string:
			bb, ok := new(big.Int).SetString(b, 0)
			if !ok {
				goto err
			}
			return &StringDef{Value: StripTrailingZeros(new(big.Int).Mod(
				new(big.Int).SetInt64(a), bb).String())}
		}
	case string:
		switch b := x2.(type) {
		case int64:
			ab, ok := new(big.Int).SetString(a, 0)
			if !ok {
				goto err
			}
			return &StringDef{Value: StripTrailingZeros(new(big.Int).Mod(ab,
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
			return &StringDef{Value: StripTrailingZeros(new(big.Int).Mod(ab, bb).String())}
		}
	}

err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '%%' to types '%T' and '%T'", x1, x2)
}
