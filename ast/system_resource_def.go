//
// The RXScript author
// SystemResource represent Go's data struct

package ast

import (
	"reflect"
)

var _ Node = (*SystemResource)(nil)

type SystemResource struct {
	obj reflect.Value
}

func (s *SystemResource) Any(handler RxeHandler) any {
	return s.obj.Interface()
}

func (s *SystemResource) String() string {
	return "'resource'"
}

func (s *SystemResource) Execute(handler RxeHandler) any {
	return s.obj.Interface()
}

func (s *SystemResource) Invoke(e RxeHandler, args []any) any {
	callArg := make([]reflect.Value, len(args))
	for i, v := range args {
		callArg[i] = reflect.ValueOf(rawValue(v, e))
	}
	returnValues := s.obj.Call(callArg)

	values := make([]Node, 0, len(returnValues))
	for _, v := range returnValues {
		values = append(values, &SystemResource{v})
	}

	if len(values) == 1 {
		return values[0]
	}

	return &ArrayDef{Value: values}
}

func (s *SystemResource) Type() CmdType {
	return SystemResourceAstCmd
}
