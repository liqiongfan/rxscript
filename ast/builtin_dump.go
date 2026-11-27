//
// The RXScript author
// built-in functions

package ast

import "fmt"

type Dump struct {
	FuncDef
}

func (p *Dump) String() string {
	return "'dump'"
}

func NewDump() *Dump {
	return &Dump{}
}

func (p *Dump) Execute(rxe RxeHandler) any {
	rawData := make([]any, 0, len(p.Args))
	for _, arg := range p.Args {
		rawData = append(rawData, NodeStringValue(arg, rxe))
	}
	fmt.Println(rawData...)
	return nil
}
