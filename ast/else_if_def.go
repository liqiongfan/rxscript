//
// The RXScript author
// ElseDef represent else if statements

package ast

var _ Node = (*ElseIfDef)(nil)

type ElseIfDef struct {
	Condition []any
	Stmts     []*FunctionBodyDef
}

func (x *ElseIfDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ElseIfDef) String() string {
	return "'else if'"
}

func (x *ElseIfDef) Type() CmdType {
	return ElseIfDefAstCmd
}

func (x *ElseIfDef) Execute(e RxeHandler) any {
	for index, cond := range x.Condition {
		if convertToBool(cond, e) {
			// execute if branch
			ret := CheckSystemError(x.Stmts[index].Execute(e), e)
			switch ret.(type) {
			case *ReturnDef, *BreakDef:
				return ret
			}
			// return true, else if branch has been run
			return &BoolDef{Value: true}
		}
	}

	return &BoolDef{Value: false}
}
