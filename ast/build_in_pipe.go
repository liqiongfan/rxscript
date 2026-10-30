//
// The RXScript author
// built-in functions

package ast

import "os"

type Pipe struct {
	FuncDef
}

func (p *Pipe) String() string {
	return "'pipe'"
}

func NewPipe() *Pipe {
	return &Pipe{}
}

func (p *Pipe) Execute(rxe RxeHandler) any {
	read, write, err := os.Pipe()
	if err != nil {
		return ThrowSystemError(err)
	}
	return &ArrayDef{Value: []Node{
		&OsFile{Value: read},
		&OsFile{Value: write},
	}}
}
