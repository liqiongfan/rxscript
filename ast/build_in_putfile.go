//
// The RXScript author
// built-in functions

package ast

import (
	"os"
	"strconv"
)

type PutFile struct {
	FuncDef
}

func (p *PutFile) String() string {
	return "'putFile'"
}

func NewPutFile() *PutFile {
	return &PutFile{}
}

// Execute putFile(filePath, fileData, perm, isAppend)
func (p *PutFile) Execute(rxe RxeHandler) any {
	var err error

	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: putFile: need at least 2 argument for file path and file data which you write into.")
	}

	fileName := rawValue(p.Args[0], rxe)
	fileData := rawValue(p.Args[1], rxe)
	filePerm := 0644

	if len(p.Args) > 2 {
		switch fp := p.Args[2].(type) {
		case *StringDef:
			filePerm, err = strconv.Atoi(fp.Value)
			if err != nil {
				return ThrowSystemError(err)
			}
		case *IntDef:
			filePerm = int(fp.Value)
		default:
			return ThrowSystemError("RXScript: putFile: filePerm must be either string or int")
		}
	}

	var flag = os.O_WRONLY | os.O_CREATE

	if len(p.Args) > 3 {
		switch isAppend := p.Args[3].(type) {
		case *BoolDef:
			if isAppend.Value {
				flag |= os.O_APPEND
			} else {
				flag |= os.O_TRUNC
			}
		case *IntDef:
			if isAppend.Value > 0 {
				flag |= os.O_APPEND
			} else {
				flag |= os.O_TRUNC
			}
		default:
			return ThrowSystemError("RXScript: putFile: isAppend must be either boolean or integer")
		}
	}

	switch fileName.(type) {
	case *StringDef, string:
	default:
		return ThrowSystemErrorWithFormat("RXScript: putFile: fileName must be string, not '%T'", fileName)
	}

	switch fileData.(type) {
	case *StringDef, string:
	default:
		return ThrowSystemErrorWithFormat("RXScript: putFile: fileData must be string, not '%T'", fileData)
	}

	f, err := os.OpenFile(fileName.(string), flag, os.FileMode(filePerm))
	if err != nil {
		return ThrowSystemError(err)
	}

	_, err = f.Write([]byte(fileData.(string)))
	if err != nil {
		return ThrowSystemError(err)
	}
	if err1 := f.Close(); err1 != nil && err == nil {
		return ThrowSystemError(err1)
	}

	return &BoolDef{true}
}
