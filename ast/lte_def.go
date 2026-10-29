//
// The RXScript author
// LteDef represent <=

package ast

import "math/big"

var _ Node = (*LteDef)(nil)

type LteDef struct {
	A Node
	B Node
}

func (x *LteDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *LteDef) String() string {
	return "'<='"
}

func (x *LteDef) Type() CmdType {
	return LteDefAstCmd
}

func (x *LteDef) Execute(e RxeHandler) any {
	x1 := rawValue(x.A.Execute(e), e)
	x2 := rawValue(x.B.Execute(e), e)

	switch a := x1.(type) {
	case int64:
		switch b := x2.(type) {
		case int64:
			return &BoolDef{Value: a <= b}
		case float64:
			return &BoolDef{Value: float64(a) <= b}
		case string:
			bb, ok := new(big.Rat).SetString(b)
			if !ok {
				goto err
			}
			if new(big.Rat).SetInt64(a).Cmp(bb) <= 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		}
	case float64:
		switch b := x2.(type) {
		case int64:
			return &BoolDef{Value: a <= float64(b)}
		case float64:
			return &BoolDef{Value: a <= b}
		case string:
			bb, ok := new(big.Rat).SetString(b)
			if !ok {
				goto err
			}
			if new(big.Rat).SetFloat64(a).Cmp(bb) <= 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		}
	case string:
		switch b := x2.(type) {
		case string:
			bb, ok := new(big.Rat).SetString(b)
			if !ok {
				return &BoolDef{Value: a <= b}
			}
			aa, ok := new(big.Rat).SetString(a)
			if !ok {
				return &BoolDef{Value: a <= b}
			}
			if aa.Cmp(bb) <= 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case float64:
			aa, ok := new(big.Rat).SetString(a)
			if !ok {
				goto err
			}
			if aa.Cmp(new(big.Rat).SetFloat64(b)) <= 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		case int64:
			aa, ok := new(big.Rat).SetString(a)
			if !ok {
				goto err
			}
			if aa.Cmp(new(big.Rat).SetInt64(b)) <= 0 {
				return &BoolDef{Value: true}
			}
			return &BoolDef{Value: false}
		}
	}
err:
	return ThrowSystemErrorWithFormat("RXScript: the two types: '%T' <= '%T' are incomparable", x1, x2)
}
