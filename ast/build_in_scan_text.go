//
// The RXScript author
// built-in functions

package ast

type ScanText struct {
	FuncDef
}

func (p *ScanText) String() string {
	return "'scanText'"
}

func NewScanText() *ScanText {
	return &ScanText{}
}

func (p *ScanText) Execute(rxe RxeHandler) any {
	if len(p.Args) < 1 {
		return ThrowSystemError("RXScript: scanNext need scanner to read data.")
	}
	dest := p.Args[0]

	switch d := dest.(type) {
	case *Scanner:
		return &StringDef{Value: d.Value.Text()}
	default:
		return ThrowSystemErrorWithFormat("RXScript: scanNext need scanner to read data.")
	}
}
