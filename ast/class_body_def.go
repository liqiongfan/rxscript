//
// The RXScript author
// ClassBodyDef represent class's attribute and methods

package ast

import (
	"strings"
)

var _ Node = (*ClassBodyDef)(nil)

// ClassBodyDef class's methods or attributes, each method or attribute
// will have the ClassBodyDef value.
type ClassBodyDef struct {
	IsMethod  bool
	IsPrivate bool
	IsPublic  bool

	Name  string
	Args  []string
	Body  *FunctionBodyDef
	Value Node

	// if is object, owner will be the object of the class
	// holds the latest class object.
	owner *Class
}

func (x *ClassBodyDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ClassBodyDef) String() string {
	if x.IsMethod {
		return "func('" + x.Name + "')"
	}
	return x.Name
}

func (x *ClassBodyDef) Type() CmdType {
	return ClassBodyDefAstCmd
}

func (x *ClassBodyDef) Execute(e RxeHandler) any {
	return x
}

func (x *ClassBodyDef) Invoke(args []any, e RxeHandler) any {
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

	hasOmitMore := false
	for _, arg := range x.Args {
		if arg == `...` {
			hasOmitMore = true
			break
		}
	}

	//args check
	if (len(x.Args) > len(args) && !hasOmitMore) || (hasOmitMore && len(x.Args)-1 > len(args)) {
		return ThrowSystemErrorWithFormat(
			"RXScript: too few arguments for %s.%s, arguments need: %s",
			x.owner.Name,
			x.Name,
			strings.Join(x.Args, ", "),
		)
	}

	// invoke method with args
	callArgs := NewNodeStore()
	callArgs.Set("this", e.GetCallingObject())
	if e.GetCallingObject() == nil && x.owner != nil {
		callArgs.Set("this", x.owner)
	}

	xArgsNum := len(x.Args)
	for i := 0; i < xArgsNum; i++ {
		arg := x.Args[i]
		if arg == `...` {
			previousArg := x.Args[i-1]
			arrayValue := &ArrayDef{Value: make([]Node, 0, xArgsNum-i+1)}
			for k := i - 1; k < len(args); k++ {
				arrayValue.Value = append(arrayValue.Value, args[k].(Node))
				i++
			}
			callArgs.Set(previousArg, arrayValue)
		} else {
			callArgs.Set(arg, args[i].(Node))
		}
	}

	e.PushArgStack(callArgs)

	// invoke
	ret := x.Body.Execute(e)
	switch x := ret.(type) {
	case *ReturnDef:
		ret = RawAstNode(x, e)
	}

	switch ret.(type) {
	case *BreakDef:
		return ThrowSystemErrorWithFormat("RXScript: break can only be used in for statements")
	}

	e.PopArgStack()

	return ret
}
