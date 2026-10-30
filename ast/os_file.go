//
// The RXScript author
// OsFile represent Go's http.Request

package ast

import (
	"os"
)

var _ Node = (*OsFile)(nil)

type OsFile struct {
	Value *os.File
}

func (x *OsFile) Any(handler RxeHandler) any {
	return x.Value
}

func (x *OsFile) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *OsFile) Execute(e RxeHandler) any {
	return x
}

func (x *OsFile) String() string {
	return "'os.File'"
}
