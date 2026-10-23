//
// The RXScript author
//

package ast

type Package struct {
	Name string

	// class definitions
	classDefs map[string]*ClassDef

	// function definitions
	funcDefs map[string]*FunctionDef
}

func (pkg *Package) GetClassDefByName(name string) (*ClassDef, bool) {
	v, ok := pkg.classDefs[name]
	return v, ok
}

func (pkg *Package) GetFuncDefByName(name string) (*FunctionDef, bool) {
	v, ok := pkg.funcDefs[name]
	return v, ok
}
