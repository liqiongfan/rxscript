//
// The RXScript author
// built-in functions

package ast

type ScanNext struct {
	FuncDef
}

func (p *ScanNext) String() string {
	return "'scanNext'"
}

func NewScanNext() *ScanNext {
	return &ScanNext{}
}

func (p *ScanNext) Execute(rxe RxeHandler) any {
	if len(p.Args) < 1 {
		return ThrowSystemError("RXScript: scanNext need scanner to read data.")
	}
	dest := p.Args[0]

	switch d := dest.(type) {
	case *Scanner:
		return &BoolDef{Value: d.Value.Scan()}
	default:
		return ThrowSystemErrorWithFormat("RXScript: scanNext need scanner to read data.")
	}
}
