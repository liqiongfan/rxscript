//
// The RXScript author
// built-in functions

package ast

type StopTicker struct {
	FuncDef
}

func (p *StopTicker) String() string {
	return "'stopTicker'"
}

func NewStopTicker() *StopTicker {
	return &StopTicker{}
}

func (p *StopTicker) Execute(rxe RxeHandler) any {
	if len(p.Args) < 1 {
		return ThrowSystemError("RXScript: need one arguments")
	}
	va := p.Args[0]

	switch v := va.(type) {
	case *TimeTicker:
		v.Value.Stop()
		return v
	default:
		return ThrowSystemErrorWithFormat("RXScript: stopTicker: invalid type '%s'", v)
	}
}
