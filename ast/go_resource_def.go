//
// The RXScript author
// GoResource represent Go's data struct

package ast

var _ Node = (*GoResource)(nil)

type GoResource struct {
	obj any
}

func (s *GoResource) Any(handler RxeHandler) any {
	return s.obj
}

func (s *GoResource) String() string {
	return "object(Go)"
}

func (s *GoResource) Execute(handler RxeHandler) any {
	return s
}

func (s *GoResource) Type() CmdType {
	return SystemResourceAstCmd
}
