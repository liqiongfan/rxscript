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

		// fetch from class object.
		ret, isPublic, isDynamic := fetchFromClassInstance(e, obj, x.Attribute)

		// check return value throw error.
		ret = CheckSystemError(ret, e)

		// fetch current class scope.
		currentClassDef := e.GetCurrentClassDef()

		// public method or attr. return direct.
		if isPublic {
			return ret
		}

		// if is dynamic or operation in class self.
		// otherwise check attr. is existed, if not exist, throw error.
		if !isDynamic &&
			(currentClassDef != nil && obj.Name != currentClassDef.Name) {
			var isDefined = false
			for _, defineOk := range currentClassDef.Body {
				if defineOk.Name == x.Attribute {
					isDefined = true
				}
			}
			if !isDefined {
				return ThrowSystemErrorWithFormat(
					"RXScript: attribute '%s' not defined in class '%s'",
					x.Attribute,
					currentClassDef.Name,
				)
			}
		}

		// if got from this, return direct.
		switch raw := x.Object.(type) {
		case *VariableDef:
			if raw.Name == ThisVar {
				return ret
			}
		}

		// otherwise return error.
		return ThrowSystemErrorWithFormat("RXScript: '%s.%s' is private, can't access out of class", obj.Name, x.Attribute)
	case *ClassDef:
		v, _ := fetchFromClassDef(e, obj, x.Attribute)
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
		return ThrowSystemErrorWithFormat(
			"RXScript: the attribute '%s' is not defined in '%s'",
			x.Attribute,
			x.Object,
		)
	}
}
