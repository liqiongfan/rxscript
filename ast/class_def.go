//
// The RXScript author
// ClassDef represent a class in RXScript

package ast

import "strings"

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

	newArgs := make([]any, 0, len(args))
	for _, arg := range args {
		switch arg := arg.(type) {
		case *SplitArrayDef:
			explodeArgs := CheckSystemError(arg.Execute(e), e)
			for _, val := range explodeArgs.([]Node) {
				newArgs = append(newArgs, RawAstNode(val, e))
			}
		default:
			newArgs = append(newArgs, RawAstNode(arg, e))
		}
	}
	args = newArgs

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

	// calling constructor if exists and is public
	for _, stmt := range x.Body {
		if !stmt.IsMethod {
			continue
		}
		if stmt.IsPublic && stmt.Name == x.Name {

			e.SetCallingObject(ci)

			hasOmitMore := false
			for _, arg := range stmt.Args {
				if arg == OmitMore {
					hasOmitMore = true
					break
				}
			}

			//args check
			if (len(stmt.Args) > len(args) && !hasOmitMore) || (hasOmitMore && len(stmt.Args)-1 > len(args)) {
				CheckSystemError(ThrowSystemErrorWithFormat(
					"RXScript: too few arguments for %s.%s, arguments need: %s",
					x.owner.Name,
					x.Name,
					strings.Join(stmt.Args, ", "),
				), e)
			}

			// call constructor function
			// prepare calling args
			callArgs := NewNodeStore()
			callArgs.Set(ThisVar, ci)

			xArgsNum := len(stmt.Args)
			realArgsNum := len(args)
			for i := 0; i < xArgsNum; i++ {
				arg := stmt.Args[i]
				if arg == OmitMore {
					previousArg := stmt.Args[i-1]
					arrayValue := &ArrayDef{Value: make([]Node, 0, xArgsNum-i+1)}
					for k := i - 1; k <= realArgsNum-1; k++ {
						arrayValue.Value = append(arrayValue.Value, args[k].(Node))
						i++
					}
					callArgs.Set(previousArg, arrayValue)
				} else {
					callArgs.Set(arg, args[i].(Node))
				}
			}
			e.PushArgStack(callArgs)

			e.PushClassDef(x)

			ret := stmt.Body.Execute(e)
			switch ret.(type) {
			case *BreakDef:
				CheckSystemError(ThrowSystemErrorWithFormat(
					"RXScript: break can only be used in for statements"), e)
			}

			e.PopClassDef()

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
