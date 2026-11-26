//
// The RXScript author
// GetAttributeDef represent any attribute obtain.

package ast

import "fmt"

var _ Node = (*GetAttributeDef)(nil)

type GetAttributeDef struct {
	Object    Node
	Attribute Node
}

func (x *GetAttributeDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *GetAttributeDef) String() string {
	return x.Object.String() + "." + x.Attribute.String()
}

func (x *GetAttributeDef) Type() CmdType {
	return GetAttributeDefAstCmd
}

func (x *GetAttributeDef) Execute(e RxeHandler) any {

	obj := CheckSystemError(x.Object.Execute(e), e)
	attr := CheckSystemError(rawValue(x.Attribute.Execute(e), e), e)
	attrName := fmt.Sprintf(`%v`, attr)

	switch obj := obj.(type) {
	case *Class:
		e.SetCallingObject(obj)

		// fetch from class object.
		ret, isPublic, isDynamic := fetchFromClassInstance(e, obj, attrName)

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
				if defineOk.Name == attrName {
					isDefined = true
				}
			}
			if !isDefined {
				return ThrowSystemErrorWithFormat(
					"RXScript: attribute '%s' not defined in class '%s'",
					attrName,
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
		return ThrowSystemErrorWithFormat("RXScript: '%s.%s' is private, can't access out of class", obj.Name, attrName)
	case *ClassDef:
		v, _ := fetchFromClassDef(e, obj, attrName)
		return v
	case *Package:
		return fetchFromPackage(obj, attrName)
	case *SystemResource:
		m := obj.obj.MethodByName(attrName)
		if m.IsValid() {
			return &SystemResource{obj: obj.obj.MethodByName(attrName)}
		}

		fd := obj.obj.FieldByName(attrName)
		if fd.IsValid() {
			return obj.toNode(fd, e)
		}
		return ThrowSystemErrorWithFormat(
			"RXScript: the attribute '%s' is not defined in '%s'",
			attrName,
			x.Object,
		)
	case *DictionaryDef:
		if v, ok := obj.Values[attrName]; ok {
			return v
		}
		obj.Values[attrName] = &DictionaryDef{Values: map[string]Node{}}
		return obj.Values[attrName]
	default:
		return ThrowSystemErrorWithFormat(
			"RXScript: the attribute '%s' is not defined in '%s'",
			attrName,
			x.Object,
		)
	}
}
