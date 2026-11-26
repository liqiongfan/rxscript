//
// The RXScript author
// BitNotDef represent ~20

package ast

import (
	"math/big"
)

var _ Node = (*BitNotDef)(nil)

type BitNotDef struct {
	A Node
}

func (x *BitNotDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *BitNotDef) String() string {
	return "'~'"
}

func (x *BitNotDef) Type() CmdType {
	return BitNotDefAstCmd
}

func (x *BitNotDef) Execute(e RxeHandler) any {
	a := rawValue(x.A.Execute(e), e)

	switch a.(type) {
	case int64:
		return &IntDef{Value: ^a.(int64)}
	case string:
		ab, ok := new(big.Int).SetString(a.(string), 10)
		if !ok {
			goto err
		}
		return &IntDef{Value: ^ab.Int64()}
	}
err:
	return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '~' to types '%T'", a)
}
