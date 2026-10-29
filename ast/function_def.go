//
// The RXScript author
// FunctionDef represent any function call, such as b(1, 20, false)

package ast

import (
	"strings"
)

var _ Node = (*FunctionDef)(nil)

// FunctionDef define a function out of the class
type FunctionDef struct {
	Name string
	Args []string
	Body *FunctionBodyDef
}

func (x *FunctionDef) Any(handler RxeHandler) any {
	return nil
}

func (x *FunctionDef) String() string {
	return "func(" + x.Name + ")"
}

func (x *FunctionDef) Type() CmdType {
	return FunctionDefAstCmd
}

func (x *FunctionDef) Execute(e RxeHandler) any {
	return x
}

// Invoke function
func (x *FunctionDef) Invoke(args []any, e RxeHandler) any {
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
	inputArgsNum := len(args)
	methodArgsNum := len(x.Args)

	hasOmitMore := false
	for _, arg := range x.Args {
		if arg == OmitMore {
			hasOmitMore = true
			break
		}
	}

	// args check
	if (methodArgsNum > inputArgsNum && !hasOmitMore) ||
		(hasOmitMore && methodArgsNum-1 > inputArgsNum) {
		return ThrowSystemErrorWithFormat(
			"RXScript: too few arguments for method %s, arguments need: %s",
			x.Name,
			strings.Join(x.Args, ", "),
		)
	}

	// invoke method with args
	callArgs := NewNodeStore()
	for i := 0; i < methodArgsNum; i++ {
		arg := x.Args[i]
		if arg == OmitMore {
			previousArg := x.Args[i-1]
			arrayValue := &ArrayDef{Value: make([]Node, 0, methodArgsNum-i+1)}
			for k := i - 1; k <= inputArgsNum-1; k++ {
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
	ret := CheckSystemError(x.Body.Execute(e), e)
	switch x := ret.(type) {
	case *ReturnDef:
		ret = CheckSystemError(x.Execute(e), e)
	}

	switch ret.(type) {
	case *BreakDef:
		return ThrowSystemErrorWithFormat("RXScript: break can only be used in for statements")
	}

	// pop arg stack
	e.PopArgStack()

	return ret
}
