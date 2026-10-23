//
// The RXScript author
// built-in functions

package ast

import "os"

type RemoveFile struct {
	FuncDef
}

func (p *RemoveFile) String() string {
	return "'removeFile'"
}

func NewRemoveFile() *RemoveFile {
	return &RemoveFile{}
}

func (p *RemoveFile) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: delFile: need at least 1 argument for file name & path.")
	}

	fileName := rawValue(p.Args[0], rxe)

	switch fileName.(type) {
	case *StringDef, string:
	default:
		return ThrowSystemErrorWithFormat("RXScript: delFile: fileName must be string, not '%T'", fileName)
	}

	err := os.Remove(fileName.(string))
	if err != nil {
		return ThrowSystemError(err)
	}

	return &BoolDef{true}
}
