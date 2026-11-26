//
// The RXScript author
// ForCondDef represent for {}.

package ast

var _ Node = (*ForCondDef)(nil)

type ForCondDef struct {
	Condition Node
	Stmts     *FunctionBodyDef
}

func (x *ForCondDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ForCondDef) String() string {
	return "object(for)"
}

// Type ForCondDefAstCmd
func (x *ForCondDef) Type() CmdType {
	return ForCondDefAstCmd
}

// Execute return the raw value of the array list
// user should use `CheckSystemError` to check value or system error occurred.
func (x *ForCondDef) Execute(e RxeHandler) any {
breakFor:
	for {
		cond := x.Condition.Execute(e)
		if !convertToBool(cond, e) {
			break
		}
		switch n := x.Stmts.Execute(e).(type) {
		case *ContinueDef:
			continue
		case *ReturnDef:
			return n
		case *BreakDef:
			break breakFor
		}
	}
	return &NullDef{}
}
