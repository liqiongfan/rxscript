//
// The RXScript author
// MultiplyDef represent *

package ast

import (
	"math/big"
)

var _ Node = (*MultiplyDef)(nil)

type MultiplyDef struct {
	A Node
	B Node
}

func (x *MultiplyDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *MultiplyDef) String() string {
	return "'*'"
}

func (x *MultiplyDef) Type() CmdType {
	return MultiplyDefAstCmd
}

func (x *MultiplyDef) Execute(e RxeHandler) any {
	x1 := rawValue(x.A.Execute(e), e)
	x2 := rawValue(x.B.Execute(e), e)

	switch a := x1.(type) {
	case int64:
		switch b := x2.(type) {
		case int64:
			return &IntDef{Value: a * b}
		case float64:
			return &DoubleDef{Value: float64(a) * b}
		}
	case float64:
		switch b := x2.(type) {
		case int64:
			return &DoubleDef{Value: a * float64(b)}
		case float64:
			return &DoubleDef{Value: a * b}
		}
	case string:
		switch b := x2.(type) {
		case string:
			ab, ok := new(big.Int).SetString(a, 10)
			if !ok {
				goto err
			}
			bb, ok := new(big.Int).SetString(b, 10)
			if !ok {
				goto err
			}
			return &StringDef{Value: new(big.Int).Mul(ab, bb).String()}
		}
	}
err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '*' to types '%T' and '%T'", x1, x2)
}
