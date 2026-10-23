//
// The RXScript author
// AndDef represent && operator

package ast

var _ Node = (*AndDef)(nil)

// AndDef && operator
// if A and B both true, the result of AndDef will be BoolDef{true}
// other than BoolDef{false}
type AndDef struct {
	A Node
	B Node
}

func (x *AndDef) Any(handler RxeHandler) any {
	return nil
}

func (x *AndDef) String() string {
	return `'&&'`
}

// Type AndDefAstCmd
func (x *AndDef) Type() CmdType {
	return AndDefAstCmd
}

// Execute the AndDef command and return the result
// must be BoolDef
// user should use `CheckSystemError` to check value or system error occurred.
func (x *AndDef) Execute(e RxeHandler) any {
	return &BoolDef{Value: convertToBool(x.A, e.(*Rxe)) &&
		convertToBool(x.B, e.(*Rxe))}
}
