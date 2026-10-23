//
// The RXScript author
// GoDef represent go {}.

package ast

var _ Node = (*GoDef)(nil)

type GoDef struct {
	Stmt *FunctionBodyDef
}

func (x *GoDef) Any(handler RxeHandler) any {
	return nil
}

func (x *GoDef) String() string {
	return "'go {}'"
}

func (x *GoDef) Type() CmdType {
	return GoDefAstCmd
}

func (x *GoDef) Execute(e RxeHandler) any {
	e.GoroutineChan() <- struct{}{}
	go func(ne RxeHandler) {
		defer func() {
			<-e.GoroutineChan()
		}()
		x.Stmt.Execute(ne)
	}(e.DumpRxe())
	return x
}
