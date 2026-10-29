//
// The RXScript author
// Class represent class object instance

package ast

var _ Node = (*Class)(nil)

type Class struct {
	// class name
	Name string

	// class variables
	// key was variable name, value is ast.AstNode
	// after constructor invoked, the variables will be the value which you need
	variables map[string]Node

	// class definitions
	def *ClassDef
}

func (x *Class) Any(handler RxeHandler) any {
	ret := make(map[string]any, len(x.variables))
	for k, v := range x.variables {
		ret[k] = v.Any(handler)
	}
	return ret
}

// GetVariables Get variables from class object
func (x *Class) GetVariables(name string) (Node, bool) {
	v, ok := x.variables[name]
	return v, ok
}

// GetClassDef Get class definition of the class object
func (x *Class) GetClassDef() *ClassDef {
	return x.def
}

// GetClassBodyDefByName Get method or attribute with the name
func (x *Class) GetClassBodyDefByName(e RxeHandler, name string) *ClassBodyDef {

	currentClassDef := e.GetCurrentClassDef()

	if currentClassDef != nil && currentClassDef.Name != x.def.Name {
		// fetch from get current class def
		for _, stmt := range currentClassDef.Body {
			if stmt.Name == name {
				return stmt
			}
		}
	}

	// fetch from the current class
	for _, stmt := range x.def.Body {
		if stmt.Name == name {
			return stmt
		}
	}

	return nil
}

func (x *Class) Type() CmdType {
	return ClassAstCmd
}

func (x *Class) Execute(e RxeHandler) any {

	return x
}

func (x *Class) String() string {
	return "object(" + x.def.Name + ")"
}
