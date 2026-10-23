//
// The RXScript author
// ForDef represent for {}.

package ast

var _ Node = (*ForDef)(nil)

type ForDef struct {
	Stmts *FunctionBodyDef
}

func (x *ForDef) Any(handler RxeHandler) any {
	return nil
}

func (x *ForDef) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (x *ForDef) String() string {
	return "'for'"
}

// Type ForRangeDefAstCmd
func (x *ForDef) Type() CmdType {
	return ForDefAstCmd
}

// Execute return the raw value of the array list
// user should use `CheckSystemError` to check value or system error occurred.
func (x *ForDef) Execute(e RxeHandler) any {
breakFor:
	for {
		switch n := x.Stmts.Execute(e).(type) {
		case *ReturnDef:
			return n
		case *BreakDef:
			break breakFor
		}
	}
	return &NullDef{}
}
