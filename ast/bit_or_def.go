//
// The RXScript author
// BitOrDef represent a | 20

package ast

import (
	"math/big"
)

var _ Node = (*BitOrDef)(nil)

type BitOrDef struct {
	A Node // it's result must be *IntDef
	B Node
}

func (x *BitOrDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *BitOrDef) String() string {
	return "'|'"
}

func (x *BitOrDef) Type() CmdType {
	return BitOrDefAstCmd
}

func (x *BitOrDef) Execute(e RxeHandler) any {
	a := rawValue(x.A.Execute(e), e)
	b := rawValue(x.B.Execute(e), e)

	switch a.(type) {
	case int64:
		switch b.(type) {
		case int64:
			return &IntDef{Value: a.(int64) | b.(int64)}
		case string:
			bb, ok := new(big.Int).SetString(b.(string), 10)
			if !ok {
				goto err
			}
			return &IntDef{Value: a.(int64) | bb.Int64()}
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
			return &IntDef{Value: ab.Int64() | bb.Int64()}
		case int64:
			ab, ok := new(big.Int).SetString(a.(string), 10)
			if !ok {
				goto err
			}
			return &IntDef{Value: ab.Int64() | b.(int64)}
		}
	}
err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '|' to types '%T' and '%T'", a, b)
}
