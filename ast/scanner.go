//
// The RXScript author
// Scanner represent Go's http.Request

package ast

import (
	"bufio"
)

var _ Node = (*Scanner)(nil)

type Scanner struct {
	Value *bufio.Scanner
}

func (x *Scanner) Any(handler RxeHandler) any {
	return x.Value
}

func (x *Scanner) Type() CmdType {
	return GoSystemResourceDefAstCmd
}

func (x *Scanner) Execute(e RxeHandler) any {
	return x
}

func (x *Scanner) String() string {
	return "'scanner'"
}
