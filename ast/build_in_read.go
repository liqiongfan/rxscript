//
// The RXScript author
// built-in functions

package ast

type Read struct {
	FuncDef
}

func (p *Read) String() string {
	return "'read'"
}

func NewRead() *Read {
	return &Read{}
}

func (p *Read) Execute(rxe RxeHandler) any {
	if len(p.Args) < 1 {
		return ThrowSystemError("RXScript: read need dest to read data.")
	}
	dest := p.Args[0]
	length := 256
	if len(p.Args) > 1 {
		switch l := p.Args[1].(type) {
		case *IntDef:
			length = int(l.Value)
		default:
			// omit length params
		}
	}

	switch d := dest.(type) {
	case *OsFile:
		data := make([]byte, length)
		_, err := d.Value.Read(data)
		if err != nil {
			return ThrowSystemError(err)
		}
		return &StringDef{Value: string(data)}
	default:
		return ThrowSystemErrorWithFormat("RXScript: read need readable object to read data.")
	}
}
