//
// The RXScript author
// LeftShiftDef represent a << 20

package ast

import (
	"math/big"
)

var _ Node = (*LeftShiftDef)(nil)

type LeftShiftDef struct {
	A Node // it's result must be *IntDef
	B Node
}

func (x *LeftShiftDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *LeftShiftDef) String() string {
	return "'<<'"
}

func (x *LeftShiftDef) Type() CmdType {
	return LeftShiftDefAstCmd
}

func (x *LeftShiftDef) Execute(e RxeHandler) any {
	a := rawValue(x.A.Execute(e), e)
	b := rawValue(x.B.Execute(e), e)

	switch a.(type) {
	case int64:
		switch b.(type) {
		case int64:
			return &IntDef{Value: a.(int64) << b.(int64)}
		}
	case string:
		switch b.(type) {
		case string:
			ab, ok := new(big.Int).SetString(a.(string), 10)
			if !ok {
				goto err
			}
			bb, ok := new(big.Int).SetString(b.(string), 10)
			if !ok {
				goto err
			}
			return &IntDef{Value: ab.Int64() << bb.Int64()}
		}
	}
err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '<<' to types '%T' and '%T'", a, b)
}
