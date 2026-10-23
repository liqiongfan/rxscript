//
// The RXScript author
// ClassDef represent a class in RXScript

package ast

var _ Node = (*ClassDef)(nil)

// ClassDef Each class was a ClassDef
type ClassDef struct {
	Name  string
	Body  []*ClassBodyDef
	owner *Class
	Base  []string
}

func (x *ClassDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ClassDef) Type() CmdType {
	return ClassDefAstCmd
}

// NewInstance create a class object, like new Class
func (x *ClassDef) NewInstance(args []any, e RxeHandler) *Class {
	// construct class instance
	ci := &Class{
		Name:      x.Name,
		variables: make(map[string]Node, 32),
		def:       x,
	}

	// set body's owner to `ci`
	for _, c := range x.Body {
		c.owner = ci
	}

	// set class def instance to `ci`
	x.owner = ci

	// copy class default value to instance
	for _, stmt := range x.Body {
		if stmt.IsMethod {
			continue
		}
		// when constructor run, the object value need to be calculated
		ci.variables[stmt.Name] = RawAstNode(stmt.Value, e).(Node)
	}

	// calling constructor if exists and is public
	for _, stmt := range x.Body {
		if !stmt.IsMethod {
			continue
		}
		if stmt.IsPublic && stmt.Name == x.Name {

			e.SetCallingObject(ci)

			// call constructor function
			// prepare calling args
			callArgs := NewNodeStore()
			callArgs.Set("this", ci)

			xArgsNum := len(stmt.Args)
			for i := 0; i < xArgsNum; i++ {
				arg := stmt.Args[i]
				if arg == `...` {
					previousArg := stmt.Args[i-1]
					arrayValue := &ArrayDef{Value: make([]Node, 0, xArgsNum-i+1)}
					for k := i - 1; k < xArgsNum; k++ {
						arrayValue.Value = append(arrayValue.Value, args[k].(Node))
						i++
					}
					callArgs.Set(previousArg, arrayValue)
				} else {
					callArgs.Set(arg, args[i].(Node))
				}
			}
			e.PushArgStack(callArgs)

			stmt.Body.Execute(e)

			e.PopArgStack()
			return ci
		}
	}

	// if no constructor, return class
	return ci
}

func (x *ClassDef) Execute(e RxeHandler) any {

	return x
}

func (x *ClassDef) String() string {
	return "class(" + x.Name + ")"
}
