//
// The RXScript author
// MinusDoubleDef represent --

package ast

var _ Node = (*MinusDoubleDef)(nil)

type MinusDoubleDef struct {
	A Node
}

func (p *MinusDoubleDef) Any(handler RxeHandler) any {
	return CheckSystemError(p.Execute(handler), handler).(Node).Any(handler)
}

func (p *MinusDoubleDef) String() string {
	return "'--'"
}

func (p *MinusDoubleDef) Execute(e RxeHandler) any {
	var ret, ov Node

	x := rawValue(p.A, e)
	switch n := x.(type) {
	case int64:
		ov = &IntDef{Value: n}
		ret = &IntDef{Value: n - 1}
	case float64:
		ov = &DoubleDef{Value: n}
		ret = &DoubleDef{Value: n - 1}
	default:
		return ThrowSystemErrorWithFormat("RXScript: cannot apply operator '--' to types '%T'", x)
	}

	switch stmt := p.A.(type) {
	case *GetAttributeDef:
		obj := CheckSystemError(stmt.Object.Execute(e), e)
		switch obj := obj.(type) {
		case *Class:
			obj.variables[stmt.Attribute] = ret
			return ov
		case *DictionaryDef:
			obj.Values[stmt.Attribute] = ret
			return ov
		default:
			return ThrowSystemErrorWithFormat("RXScript: cannot found object with name: '%s'", stmt.Object)
		}
	case *VariableDef:
		e.SetArgs(stmt.Name, ret)
		return ov
	default:
		name := rawValue(p.A, e)
		e.SetArgs(name.(string), ret)
		return ov
	}
}

func (p *MinusDoubleDef) Type() CmdType {
	return MinusDoubleDefAstCmd
}
