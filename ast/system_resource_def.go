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
	return "object(Res)"
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
		if !v.IsValid() || v.IsZero() || (v.Kind() == reflect.Pointer && v.IsNil()) {
			values = append(values, &NullDef{})
			continue
		}
		switch v.Kind() {
		case reflect.Bool:
			values = append(values, &BoolDef{Value: v.Bool()})
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values = append(values, &IntDef{Value: v.Int()})
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			values = append(values, &IntDef{Value: int64(v.Uint())})
		case reflect.String:
			values = append(values, &StringDef{Value: v.String()})
		case reflect.Float32, reflect.Float64:
			values = append(values, &DoubleDef{Value: v.Float()})
		case reflect.Map:
			values = append(values, toNodeValue(v.Interface(), e).(Node))
		case reflect.Slice:
			values = append(values, toNodeValue(v.Interface(), e).(Node))
		default:
			values = append(values, &SystemResource{obj: v})
		}
	}

	if len(values) == 1 {
		return values[0]
	}

	return &ArrayDef{Value: values}
}

func (s *SystemResource) Type() CmdType {
	return SystemResourceAstCmd
}
