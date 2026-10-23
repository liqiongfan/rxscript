//
// The RXScript author
// ElseDef represent else statements

package ast

var _ Node = (*ElseDef)(nil)

type ElseDef struct {
	Stmt *FunctionBodyDef
}

func (x *ElseDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ElseDef) String() string {
	return "'else'"
}

func (x *ElseDef) Type() CmdType {
	return ElseDefAstCmd
}

func (x *ElseDef) Execute(e RxeHandler) any {
	if x.Stmt == nil {
		return &NullDef{}
	}

	return CheckSystemError(x.Stmt.Execute(e), e)
}
