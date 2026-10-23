//
// The RXScript author
// GetNodeAttributeDef represent any attribute obtain.

package ast

var _ Node = (*GetNodeAttributeDef)(nil)

type GetNodeAttributeDef struct {
	Object    *FuncCallDef
	Attribute string
}

func (x *GetNodeAttributeDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *GetNodeAttributeDef) String() string {
	return "'" + x.Object.String() + "." + x.Attribute + "'"
}

func (x *GetNodeAttributeDef) Type() CmdType {
	return GetNodeAttributeDefAstCmd
}

func (x *GetNodeAttributeDef) Execute(e RxeHandler) any {

	objectInstance := x.Object.Execute(e)
	switch object := objectInstance.(type) {
	case *Class:
		e.SetCallingObject(object)
		ret, _ := fetchFromClassInstance(object, x.Attribute)
		switch v := ret.(type) {
		case *ClassBodyDef:
			if v.IsPrivate {
				return ThrowSystemError("RXScript: calling private method object: " + object.Name + ", method:" + x.Attribute)
			}
		}
		return ret
	case *SystemResource:
		return &SystemResource{obj: object.obj.MethodByName(x.Attribute)}
	default:
		return ThrowSystemErrorWithFormat("RXScript: the chain of func calling need return object. not: '%s'", object)
	}
}
