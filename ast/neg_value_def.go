//
// The RXScript author
// NegValueDef represent -10

package ast

import "encoding/json"

var _ Node = (*NegValueDef)(nil)

type NegValueDef struct {
	Value Node
}

func (x *NegValueDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *NegValueDef) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(x.Value)
	if err != nil {
		return nil, err
	}
	return append([]byte("-"), buf...), nil
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
	default:
		return ThrowSystemErrorWithFormat("RXScript: unknown negative type: '%s'", v)
	}
}
