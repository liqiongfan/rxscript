//
// The RXScript author
// GteDef represent >=.

package ast

var _ Node = (*GteDef)(nil)

type GteDef struct {
	A Node
	B Node
}

func (x *GteDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *GteDef) String() string {
	return "'>='"
}

func (x *GteDef) Type() CmdType {
	return GteDefAstCmd
}

func (x *GteDef) Execute(e RxeHandler) any {
	x1 := rawValue(x.A.Execute(e), e)
	x2 := rawValue(x.B.Execute(e), e)

	switch a := x1.(type) {
	case int64:
		switch b := x2.(type) {
		case int64:
			return &BoolDef{Value: a >= b}
		case float64:
			return &BoolDef{Value: float64(a) >= b}
		}
	case float64:
		switch b := x2.(type) {
		case int64:
			return &BoolDef{Value: a >= float64(b)}
		case float64:
			return &BoolDef{Value: a >= b}
		}
	case string:
		switch b := x2.(type) {
		case string:
			return &BoolDef{Value: a >= b}
		}
	}

	return ThrowSystemErrorWithFormat("RXScript: the two types: '%T' >= '%T' are incomparable", x1, x2)
}
