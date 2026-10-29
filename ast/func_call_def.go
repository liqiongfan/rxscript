//
// The RXScript author
// FuncCallDef represent a() or a.b() or a()() && a.b()() && a.b.c.d() and so on ...

package ast

var _ Node = (*FuncCallDef)(nil)

type FuncCallDef struct {
	Name any

	// real arg values.
	Args []any

	// 1: name.name AstNode
	// 2: name  string
	// 3. result's of above 1 & 2.
	// deprecated
	NameType int
}

func (x *FuncCallDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *FuncCallDef) String() string {
	return x.Name.(Node).String()
}

func (x *FuncCallDef) Type() CmdType {
	return FuncCallDefAstCmd
}

func (x *FuncCallDef) Execute(e RxeHandler) any {
	obj := x.Name

	for {
		switch node := obj.(type) {
		case *ClassDef:
			return node.NewInstance(x.Args, e)
		case *ClassBodyDef:
			if node.IsMethod {
				return CheckSystemError(node.Invoke(x.Args, e), e)
			}
			return ThrowSystemErrorWithFormat("RXScript: object(%s) is not a callable", x.Name)
		case *FunctionDef:
			return CheckSystemError(node.Invoke(x.Args, e), e)
		case *SystemErrorDef:
			return CheckSystemError(node.Execute(e), e)
		case *VariableDef:
			obj = CheckSystemError(node.Execute(e), e)
		case *SystemResource:
			return CheckSystemError(node.Invoke(e, x.Args), e)
		case interface {
			WithArgs(nd any, e RxeHandler, args ...any)
			Execute(RxeHandler) any
		}:
			node.WithArgs(node, e, x.Args...)
			return CheckSystemError(node.Execute(e), e)
		case *GetAttributeDef:
			obj = CheckSystemError(node.Execute(e), e)
		case *DictionaryDef, *Class, *ArrayDef, *TimeTimer, *TimeTicker, *ReturnDef,
			*IntDef, *StringDef, *BoolDef, *DoubleDef, *NullDef, *ChannelDef, *WaitDef, *HttpRequest:
			return ThrowSystemErrorWithFormat("RXScript: object(%s) is not a callable", x.Name)
		case Node:
			obj = CheckSystemError(node.Execute(e), e)
		default:
			panic(node)
		}
	}
}
