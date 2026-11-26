//
// The RXScript author
// VariableAssignDef represent b = xx

package ast

import "fmt"

var _ Node = (*VariableAssignDef)(nil)

type VariableAssignDef struct {
	Name  Node
	Value Node
}

func (x *VariableAssignDef) Any(handler RxeHandler) any {
	return `{}`
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
		attr := CheckSystemError(rawValue(a.Attribute.Execute(e), e), e)
		attrName := fmt.Sprintf(`%v`, attr)
		switch obj := obj.(type) {
		case *Class:
			if _, ok := obj.variables[attrName]; ok {
				for _, stmt := range obj.def.Body {
					if stmt.Name == attrName && stmt.IsPrivate {
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
			obj.variables[attrName] = value.(Node)
		case *DictionaryDef:
			obj.Values[attrName] = value.(Node)
		default:
			panic("RXScript: Unsupported feature")
		}
	case *ArrayIndexDef:
		node := a.Value.Execute(e)
		switch n := node.(type) {
		case *ArrayDef:
			v := n.Execute(e)
			if v == nil {
				return ThrowSystemError("RXScript: null value can't be indexed")
			}
			if a.Index >= int64(len(v.(*ArrayDef).Value)) {
				return ThrowSystemErrorWithFormat("RXScript: index out of range, index: %d, total: %d", a.Index, len(v.([]Node)))
			}
			v.(*ArrayDef).Value[a.Index] = x.Value.Execute(e).(Node)
		default:
			panic("RXScript: Unsupported feature")
		}
	case *DictionaryIndexDef:
		node := a.Values.Execute(e)
		switch n := node.(type) {
		case *DictionaryDef:
			key := a.Key.Execute(e)
			switch k := key.(type) {
			case *StringDef:
				n.Values[k.Value] = x.Value.Execute(e).(Node)
			default:
				panic("RXScript: Unsupported feature")
			}
		default:
			panic("RXScript: Unsupported feature")
		}
	default:
		name := rawValue(x.Name.Execute(e), e)
		e.SetArgs(name.(string), value)
	}

	return x
}
