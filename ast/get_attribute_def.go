//
// The RXScript author
// GetAttributeDef represent any attribute obtain.

package ast

var _ Node = (*GetAttributeDef)(nil)

type GetAttributeDef struct {
	Object    Node
	Attribute string
}

func (x *GetAttributeDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *GetAttributeDef) String() string {
	return x.Object.String() + "." + x.Attribute
}

func (x *GetAttributeDef) Type() CmdType {
	return GetAttributeDefAstCmd
}

func (x *GetAttributeDef) Execute(e RxeHandler) any {

	var obj = CheckSystemError(x.Object.Execute(e), e)

	switch obj := obj.(type) {
	case *Class:
		e.SetCallingObject(obj)
		ret, isPublic := fetchFromClassInstance(obj, x.Attribute)

		ret = CheckSystemError(ret, e)
		if isPublic {
			return ret
		}

		switch raw := x.Object.(type) {
		case *VariableDef:
			if raw.Name == "this" {
				return ret
			}
		}
		return ThrowSystemErrorWithFormat("RXScript: '%s.%s' is private, can't access out of class", obj.Name, x.Attribute)
	case *ClassDef:
		v, _ := fetchFromClassDef(obj, x.Attribute)
		return v
	case *Package:
		return fetchFromPackage(obj, x.Attribute)
	case *SystemResource:
		return &SystemResource{obj: obj.obj.MethodByName(x.Attribute)}
	case *DictionaryDef:
		if v, ok := obj.Values[x.Attribute]; ok {
			return v
		}
		obj.Values[x.Attribute] = &DictionaryDef{Values: map[string]Node{}}
		return obj.Values[x.Attribute]
	default:
		return ThrowSystemErrorWithFormat("RXScript: the attribute '%s' is not defined in '%s'", x.Attribute, x.Object)
	}
}
