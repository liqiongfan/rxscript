//
// The RXScript author
// built-in functions

package ast

type Channel struct {
	FuncDef
}

func (p *Channel) String() string {
	return "'channel'"
}

func NewChannel() *Channel {
	return &Channel{}
}

func (p *Channel) Execute(e RxeHandler) any {
	if len(p.Args) == 1 {
		num := p.Args[0]
		switch v := num.(type) {
		case *IntDef:
			return &ChannelDef{Value: make(chan any, v.Value)}
		default:
			return ThrowSystemErrorWithFormat("RXScript: channel: invalid argument type, not '%S'", num)
		}
	}
	return &ChannelDef{Value: make(chan any)}
}
