//
// The RXScript author
// FunctionBodyDef represent any statement in { xx }

package ast

var _ Node = (*FunctionBodyDef)(nil)

type FunctionBodyDef struct {
	Stmt  []Node
	Defer []Node
}

func (x *FunctionBodyDef) Any(handler RxeHandler) any {
	return CheckSystemError(x.Execute(handler), handler).(Node).Any(handler)
}

func (x *FunctionBodyDef) String() string {
	return "'{}'"
}

func (x *FunctionBodyDef) Type() CmdType {
	return FunctionBodyDefAstCmd
}

func (x *FunctionBodyDef) Execute(e RxeHandler) any {

	defer func() {
		// run defer
		if len(x.Defer) > 0 {
			for i := len(x.Defer) - 1; i >= 0; i-- {
				CheckSystemError(x.Defer[i].Execute(e), e)
			}
		}
	}()

	for _, stmt := range x.Stmt {
		// run statements
		ret := CheckSystemError(stmt.Execute(e), e)

		// if return
		if stmt.Type() == ReturnDefAstCmd {
			return &ReturnDef{Value: ret}
		}

		// if break
		switch ret.(type) {
		case *BreakDef:
			return ret
		}
	}

	return &NullDef{}
}
