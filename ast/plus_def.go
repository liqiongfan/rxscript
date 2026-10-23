//
// The RXScript author
// PlusDef represent +

package ast

import (
	"math/big"
)

var _ Node = (*PlusDef)(nil)

type PlusDef struct {
	A Node
	B Node
}

func (x *PlusDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *PlusDef) String() string {
	return "'+'"
}

func (x *PlusDef) Type() CmdType {
	return PlusDefAstCmd
}

func (x *PlusDef) Execute(e RxeHandler) any {
	a := rawValue(x.A.Execute(e), e)
	b := rawValue(x.B.Execute(e), e)

	switch a.(type) {
	case int64:
		switch b.(type) {
		case int64:
			return &IntDef{Value: a.(int64) + b.(int64)}
		case float64:
			return &DoubleDef{Value: float64(a.(int64)) + b.(float64)}
		}
	case float64:
		switch b.(type) {
		case int64:
			return &DoubleDef{Value: a.(float64) + float64(b.(int64))}
		case float64:
			return &DoubleDef{Value: a.(float64) + b.(float64)}
		}
	case string:
		switch b.(type) {
		case string:
			ab, ok := new(big.Int).SetString(a.(string), 10)
			if !ok {
				return &StringDef{Value: a.(string) + b.(string)}
			}
			bb, ok := new(big.Int).SetString(b.(string), 10)
			if !ok {
				return &StringDef{Value: a.(string) + b.(string)}
			}
			return &StringDef{Value: new(big.Int).Add(ab, bb).String()}
		}
	}

	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '+' to types '%T' and '%T'", a, b)
}
