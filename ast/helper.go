//
// The RXScript author
// engine helper functions to operate the raw RXScript value

package ast

import "strings"

func convertToBool(qv any, e RxeHandler) bool {
	qv = rawValue(qv, e)
	switch q := qv.(type) {
	case int64:
		return q > 0
	case float64:
		return q > 0
	case string:
		return q != ``
	case bool:
		return q
	case []any:
		return len(q) > 0
	case chan any:
		return len(q) > 0
	case map[string]any:
		return len(q) > 0
	case *SplitArrayDef:
		return len(CheckSystemError(q.Execute(e), e).([]Node)) > 0
	case nil:
		return false
	default:
		return false
	}
}

// fetchFromRxe from environment to fetch value
func fetchFromRxe(e RxeHandler, name string) any {
	// from args
	if v, ok := e.GetFromArgs(name); ok {
		return v
	}

	// from global variables
	if v, ok := e.GetFromVariables(name); ok {
		return v
	}

	// from func
	if v, ok := e.GetFromFuncDefs(name); ok {
		return v
	}

	// from class
	if v, ok := e.GetFromClassDefs(name); ok {
		// deal with inherit
		v := v.(*ClassDef)
		if v.Base != nil && len(v.Base) > 0 {
			recursiveClassBase(e, v, v.Base)
		}
		return v
	}

	// from package
	if v, ok := e.GetFromPackageDefs(name); ok {
		return v
	}

	// build-in function list
	if v := fetchBuildInFunc(name); v != nil {
		return v
	}

	// not found
	return ThrowSystemErrorWithFormat(
		"RXScript: unknown variables '%s', please define it before you use.",
		name,
	)
}

// recursiveClassBase do with the class inherit
func recursiveClassBase(e RxeHandler, base *ClassDef, baseClass []string) Node {
	bodyDef := make([]*ClassBodyDef, 0, 32)
	for _, baseName := range baseClass {
		if parentClass, ok := e.GetFromClassDefs(baseName); !ok {
			return &SystemErrorDef{Value: "RXScript: class not defined: " + baseName}
		} else {
			parent := parentClass.(*ClassDef)
			if parent.Base != nil && len(parent.Base) > 0 {
				recursiveClassBase(e, base, parent.Base)
			}
			for _, stmt := range parent.Body {
				if stmt.IsPrivate {
					continue
				}
				stmt.ClassName = parent.Name
				bodyDef = append(bodyDef, stmt)
			}
		}
	}
	base.Body = append(base.Body, bodyDef...)
	return &NullDef{}
}

// fetchFromClassInstance fetch data from object instance
func fetchFromClassInstance(e RxeHandler, ci *Class, name string) (
	any, bool /*isPublic*/, bool /* is dynamic*/) {
	if v, ok := ci.GetVariables(name); ok {
		// GetVariables is public or dynamic set
		return v, true, true
	}

	// method or attribute from object instance
	ma := ci.GetClassBodyDefByName(e, name)
	if ma != nil {
		if ma.IsMethod {
			// constructor
			if ma.Name == ma.ClassName {
				return ma, ma.IsPublic, true
			}
			// not constructor
			return ma, ma.IsPublic, false
		}
		// not method, need execute the node to fetch the value
		return RawAstNode(ma.Value, e), ma.IsPublic, false
	}

	// not from instance, from class
	v, ok := fetchFromClassDef(e, ci.GetClassDef(), name)
	return v, ok, false
}

// fetchFromClassDef fetch from class definition
func fetchFromClassDef(e RxeHandler, cd *ClassDef, name string) (any, bool) {
	className := cd.Name
	currentClassDef := e.GetCurrentClassDef()
	if currentClassDef != nil {
		className = currentClassDef.Name
	}

	// method or attribute
	for _, stmt := range cd.Body {
		if stmt.Name == name && cd.Name == stmt.ClassName {
			if stmt.IsMethod {
				// constructor
				if stmt.Name == stmt.ClassName {
					return stmt, true
				}
				// not constructor
				return stmt.Body, stmt.IsPublic
			}
			// not method
			return RawAstNode(stmt.Value, e), stmt.IsPublic
		}
	}

	return ThrowSystemErrorWithFormat(
		"RXScript: unknown attr: '%s' in class: '%s'",
		name,
		className,
	), false
}

// fetchFromPackage fetch from package
func fetchFromPackage(cd *Package, name string) any {
	if v, ok := cd.GetClassDefByName(name); ok {
		return v
	}

	if v, ok := cd.GetFuncDefByName(name); ok {
		return v
	}

	return ThrowSystemErrorWithFormat(
		"RXScript: unknown attr: '%s' in package: '%s'",
		name,
		cd.Name,
	)
}

// StripTrailingZeros trim right ending zero.
func StripTrailingZeros(s string) string {
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimRight(s, ".")
	}
	return s
}
