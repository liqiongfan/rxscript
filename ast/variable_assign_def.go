//
// The RXScript author
// VariableAssignDef represent b = xx

package ast

var _ Node = (*VariableAssignDef)(nil)

type VariableAssignDef struct {
	Name  Node
	Value Node
}

func (x *VariableAssignDef) Any(handler RxeHandler) any {
	return nil
}

func (x *VariableAssignDef) String() string {
	return x.Name.String()
}

func (x *VariableAssignDef) Type() CmdType {
	return VariableAssignDefAstCmd
}

func (x *VariableAssignDef) Execute(e RxeHandler) any {

	value := x.Value.Execute(e)

	switch a := x.Name.(type) {
	case *GetAttributeDef:
		obj := CheckSystemError(a.Object.Execute(e), e)
		switch obj := obj.(type) {
		case *Class:
			if _, ok := obj.variables[a.Attribute]; ok {
				for _, stmt := range obj.def.Body {
					if stmt.Name == a.Attribute && stmt.IsPrivate {
						switch v := a.Object.(type) {
						case *VariableDef:
							if v.Name != ThisVar {
								return ThrowSystemErrorWithFormat(
									"RXScript: %s.%s is private, can't access out of class",
									obj.Name,
									stmt.Name,
								)
							}
						}
					}
				}
			}
			obj.variables[a.Attribute] = value.(Node)
		case *DictionaryDef:
			switch v := a.Object.(type) {
			case *VariableDef:
				if v.Name != ThisVar {
					return ThrowSystemErrorWithFormat(
						"RXScript: %s is private, can't access out of class",
						x.Name.String(),
					)
				}
			}
			obj.Values[a.Attribute] = value.(Node)
		default:
			panic(a)
		}
	default:
		name := rawValue(x.Name.Execute(e), e)
		e.SetArgs(name.(string), value)
	}

	return x
}
