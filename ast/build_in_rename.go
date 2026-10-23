//
// The RXScript author
// built-in functions

package ast

import "os"

type Rename struct {
	FuncDef
}

func (p *Rename) String() string {
	return "'rename'"
}

func NewRename() *Rename {
	return &Rename{}
}

func (p *Rename) Execute(rxe RxeHandler) any {
	if len(p.Args) <= 0 {
		return ThrowSystemError("RXScript: rename: need at least 2 argument for file name & path.")
	}

	fileName := rawValue(p.Args[0], rxe)
	newFileName := rawValue(p.Args[1], rxe)

	switch fileName.(type) {
	case *StringDef, string:
	default:
		return ThrowSystemErrorWithFormat("RXScript: rename: origin file name must be string, not '%T'", fileName)
	}

	switch newFileName.(type) {
	case *StringDef, string:
	default:
		return ThrowSystemErrorWithFormat("RXScript: rename: new file name must be string, not '%T'", newFileName)
	}

	err := os.Rename(fileName.(string), newFileName.(string))
	if err != nil {
		return ThrowSystemError(err)
	}

	return &BoolDef{true}
}
