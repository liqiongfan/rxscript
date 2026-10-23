//
// The RXScript author
// built-in functions

package ast

import (
	"os"
)

type ListDir struct {
	FuncDef
}

func (p *ListDir) String() string {
	return "'listDir'"
}

func NewListDir() *ListDir {
	return &ListDir{}
}

func (p *ListDir) Execute(rxe RxeHandler) any {
	var path any = "."
	if len(p.Args) > 0 {
		path = rawValue(p.Args[0], rxe)
	}

	switch path := path.(type) {
	case string:
		dirEntry, err := os.ReadDir(path)
		if err != nil {
			return ThrowSystemError(err)
		}
		ret := &ArrayDef{Value: make([]Node, 0, len(dirEntry))}
		for _, de := range dirEntry {
			ret.Value = append(ret.Value, &StringDef{Value: de.Name()})
		}
		return ret
	default:
		return ThrowSystemErrorWithFormat(`RXScript: argument need to be a string, not '%T'`, path)
	}
}
