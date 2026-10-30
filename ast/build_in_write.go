//
// The RXScript author
// built-in functions

package ast

type Write struct {
	FuncDef
}

func (p *Write) String() string {
	return "'write'"
}

func NewWrite() *Write {
	return &Write{}
}

func (p *Write) Execute(rxe RxeHandler) any {
	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: write need dest & data to write.")
	}
	dest := p.Args[0]
	data := p.Args[1]

	var value string

	switch l := data.(type) {
	case *StringDef:
		value = l.Value
	default:
		// omit length params
	}

	switch d := dest.(type) {
	case *OsFile:
		_, err := d.Value.WriteString(value)
		if err != nil {
			return ThrowSystemError(err)
		}
		return &BoolDef{Value: true}
	default:
		return ThrowSystemErrorWithFormat("RXScript: write need writeable object to write data.")
	}
}
