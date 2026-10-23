//
// The RXScript author
// IfStmtDef represent import "xx"

package ast

import (
	"os"
	"path/filepath"
	"strings"
)

var _ Node = (*Import)(nil)

type Import struct {
	Path     string
	As       string
	IsSystem bool
}

func (i *Import) Any(handler RxeHandler) any {
	return nil
}

func (i *Import) String() string {
	return "'import'"
}

func (i *Import) Execute(e RxeHandler) any {
	var err error
	var buff []byte
	var op []os.DirEntry

	if i.IsSystem {
		op, err = stdLibrary.ReadDir(i.Path)
		if err != nil {
			return ThrowSystemError(err)
		}
	} else {
		op, err = os.ReadDir(i.Path)
		if err != nil {
			return ThrowSystemError(err)
		}
	}

	for _, d := range op {
		if d.IsDir() {
			continue
		}

		if !strings.HasSuffix(d.Name(), ScriptSuffix) {
			continue
		}

		fileName := filepath.Join(i.Path, d.Name())

		if i.IsSystem {
			buff, err = stdLibrary.ReadFile(
				strings.ReplaceAll(i.Path, `\`, `/`) + `/` + d.Name())
			if err != nil {
				return ThrowSystemError(err)
			}
		} else {
			buff, err = os.ReadFile(fileName)
			if err != nil {
				return ThrowSystemError(err)
			}
		}

		lex := &RxeLexer{
			buff:     string(buff) + "\x00",
			line:     1,
			column:   1,
			fileName: fileName,
		}

		if rxeParse(lex) == 0 {
			pack := &Package{
				Name:      i.As,
				classDefs: make(map[string]*ClassDef, 32),
				funcDefs:  make(map[string]*FunctionDef, 32),
			}
			for _, stmt := range lex.result {
				switch v := stmt.(type) {
				case *Import:
					CheckSystemError(v.Execute(e), e)
				case *ClassDef:
					pack.classDefs[v.Name] = v
				case *FunctionDef:
					if v.Name == `main` {
						return ThrowSystemError("RXScript: main can't be redefined.")
					} else {
						pack.funcDefs[v.Name] = v
					}
				}
			}
			if _, ok := e.(*Rxe).packageDefs[i.As]; ok {
				for name, v := range pack.classDefs {
					e.(*Rxe).packageDefs[i.As].classDefs[name] = v
				}
				for name, v := range pack.funcDefs {
					e.(*Rxe).packageDefs[i.As].funcDefs[name] = v
				}
			} else {
				e.(*Rxe).packageDefs[i.As] = pack
			}

		} else {
			return &SystemErrorDef{Value: fileName}
		}
	}

	return i
}

func (i *Import) Type() CmdType {
	return ImportAstCmd
}
