//
// The RXScript author
// built-in functions

package ast

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Sha256 struct {
	FuncDef
}

func (p *Sha256) String() string {
	return "'sha256'"
}

func NewSha256() *Sha256 {
	return &Sha256{}
}

func (p *Sha256) Execute(rxe RxeHandler) any {
	if len(p.Args) > 0 {
		var arg = p.Args[0]
		switch a := arg.(type) {
		case *StringDef:
			m := sha256.New()
			m.Write([]byte(a.Value))
			return &StringDef{Value: hex.EncodeToString(m.Sum(nil))}
		default:
			m := sha256.New()
			m.Write([]byte(fmt.Sprintf("%v", a)))
			return &StringDef{Value: hex.EncodeToString(m.Sum(nil))}
		}
	}
	return &NullDef{}
}
