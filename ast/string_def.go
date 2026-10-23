//
// The RXScript author
// StringDef represent "xxx"

package ast

var _ Node = (*StringDef)(nil)

type StringDef struct {
	Value string
}

func (x *StringDef) Any(handler RxeHandler) any {
	return x.Value
}

func (x *StringDef) Type() CmdType {
	return StringDefAstCmd
}

func (x *StringDef) Execute(e RxeHandler) any {
	return x
}

func (x *StringDef) String() string {
	return "string(" + x.Value + ")"
}
