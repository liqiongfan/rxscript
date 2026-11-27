//
// The RXScript author
// built-in functions

package ast

type Invoke struct {
	FuncDef
}

func (p *Invoke) String() string {
	return "'invoke'"
}

func NewInvoke() *Invoke {
	return &Invoke{}
}

func (p *Invoke) Execute(rxe RxeHandler) any {
	if len(p.Args) < 2 {
		return ThrowSystemError("RXScript: invoke need object and method.")
	}

	obj := p.Args[0]
	method := p.Args[1]

	var args []any
	if len(p.Args) > 2 {
		args = p.Args[2:]
	}

	switch o := obj.(type) {
	case *Class:
		switch m := method.(type) {
		case *StringDef:
			ret, isPublic, isDynamic := fetchFromClassInstance(rxe, o, m.Value)

			// check return value throw error.
			ret = CheckSystemError(ret, rxe)

			// fetch current class scope.
			currentClassDef := rxe.GetCurrentClassDef()

			if !isDynamic &&
				(currentClassDef != nil && o.Name != currentClassDef.Name) {
				var isDefined = false
				for _, defineOk := range currentClassDef.Body {
					if defineOk.Name == m.Value {
						isDefined = true
					}
				}
				if !isDefined {
					return ThrowSystemErrorWithFormat(
						"RXScript: attribute '%s' not defined in class '%s'",
						m.Value,
						currentClassDef.Name,
					)
				}
			}

			if isPublic {
				// call
				switch cb := ret.(type) {
				case *ClassBodyDef:
					rxe.SetCallingObject(cb.owner)
					return cb.Invoke(args, rxe)
				default:
					return ThrowSystemErrorWithFormat("RXScript: invoke need callable, not '%s'", cb)
				}
			}

			return ThrowSystemErrorWithFormat("RXScript: '%s.%s' is private, can't access out of class", o.Name, m.Value)
		}
	}

	return &NullDef{}
}
