//
// The RXScript author
// built-in functions

package ast

import "os"

type ReadFile struct {
	FuncDef
}

func (p *ReadFile) String() string {
	return "'readFile'"
}

func NewReadFile() *ReadFile {
	return &ReadFile{}
}

func (p *ReadFile) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: readFile: need at least 1 argument for file name & path.")
	}

	fileName := rawValue(p.Args[0], rxe)

	switch fileName.(type) {
	case *StringDef, string:
	default:
		return ThrowSystemErrorWithFormat("RXScript: readFile: fileName must be string, not '%T'", fileName)
	}

	fileData, err := os.ReadFile(fileName.(string))
	if err != nil {
		return ThrowSystemError(err)
	}

	return &StringDef{Value: string(fileData)}
}
