//
// The RXScript author
// built-in functions

package ast

import "fmt"

type Println struct {
	FuncDef
}

func (p *Println) String() string {
	return "'println'"
}

func NewPrintln() *Println {
	return &Println{}
}

func (p *Println) Execute(rxe RxeHandler) any {
	rawData := make([]any, 0, len(p.Args))
	for _, arg := range p.Args {
		//print(NodeStringValue(arg, rxe), " ")
		rawData = append(rawData, NodeStringValue(arg, rxe))
	}
	//println()
	fmt.Println(rawData...)
	return nil
}
