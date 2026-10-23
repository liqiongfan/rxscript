//
// The RXScript author
// IfStmtDef represent if xx {} else if xx {} else if xx {} else {}.

package ast

var _ Node = (*IfStmtDef)(nil)

type IfStmtDef struct {
	Condition Node
	Stmt      *FunctionBodyDef
	ElseIfdef *ElseIfDef
	ElseDef   *ElseDef
}

func (x *IfStmtDef) Any(handler RxeHandler) any {
	return nil
}

func (x *IfStmtDef) String() string {
	return "'if'"
}

func (x *IfStmtDef) Type() CmdType {
	return IfStmtDefAstCmd
}

func (x *IfStmtDef) Execute(e RxeHandler) any {
	// check if condition
	if convertToBool(x.Condition.Execute(e), e) {
		ret := CheckSystemError(x.Stmt.Execute(e), e)
		switch ret.(Node).Type() {
		case ReturnDefAstCmd, BreakDefAstCmd:
			return ret
		default:
		}
	} else {
		// else if condition
		if len(x.ElseIfdef.Condition) > 0 {
			elseIfResult := x.ElseIfdef.Execute(e)
			switch elseIfResult.(type) {
			case *ReturnDef, *BreakDef:
				return elseIfResult
			case *BoolDef:
				// else
				if !convertToBool(elseIfResult, e) {
					elseResult := x.ElseDef.Execute(e)
					switch elseResult.(type) {
					case *ReturnDef, *BreakDef:
						return elseResult
					}
				}
			}
		} else {
			// else if empty, direct use else
			elseResult := x.ElseDef.Execute(e)
			switch elseResult.(type) {
			case *ReturnDef, *BreakDef:
				return elseResult
			}
		}
	}

	return x
}
