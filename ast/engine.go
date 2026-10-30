//
// The RXScript author
// Engine descriptor, global executor's environment

package ast

type Rxe struct {
	// all variables
	variables NodeStore

	// function calling args
	callArgs NodeStore

	// class definitions
	classDefs map[string]*ClassDef

	// function definitions
	funcDefs map[string]*FunctionDef

	// package definitions
	packageDefs map[string]*Package

	// calling object
	callingObject *Class

	// args stack
	argsStack []NodeStore

	// goroutine number
	goroutineChan chan struct{}

	// current class definitions
	currentClassDef *ClassDef
	callingScope    []*ClassDef
}

func (e *Rxe) GetCurrentClassDef() *ClassDef {
	return e.currentClassDef
}

func (e *Rxe) SetCurrentClassDef(v *ClassDef) {
	e.currentClassDef = v
}

func (e *Rxe) PushClassDef(def *ClassDef) {
	e.currentClassDef = def
	e.callingScope = append(e.callingScope, def)
}

func (e *Rxe) PopClassDef() *ClassDef {
	e.callingScope = e.callingScope[:len(e.callingScope)-1]

	lastIndex := len(e.callingScope) - 1
	e.currentClassDef = nil
	if lastIndex >= 0 {
		e.currentClassDef = e.callingScope[lastIndex]
	}

	return e.currentClassDef
}

func (e *Rxe) GoroutineChan() chan struct{} {
	return e.goroutineChan
}

func (e *Rxe) SetCallingObject(c *Class) {
	e.callingObject = c
}

func (e *Rxe) GetCallingObject() *Class {
	return e.callingObject
}

func (e *Rxe) PushArgStack(arg NodeStore) {
	e.argsStack = append(e.argsStack, e.callArgs, arg)
	e.callArgs = arg
}

func (e *Rxe) PopArgStack() {
	e.callingObject = nil
	e.callArgs = e.argsStack[len(e.argsStack)-2]
	e.argsStack = e.argsStack[:len(e.argsStack)-2]
}

func (e *Rxe) GetVariables() NodeStore {
	return e.variables
}

func (e *Rxe) GetCallArgs() NodeStore {
	return e.callArgs
}

func (e *Rxe) GetFromArgs(name string) (any, bool) {
	v, ok := e.callArgs[name]
	return v, ok
}

func (e *Rxe) GetFromVariables(name string) (any, bool) {
	v, ok := e.variables[name]
	return v, ok
}

func (e *Rxe) GetFromFuncDefs(name string) (any, bool) {
	v, ok := e.funcDefs[name]
	return v, ok
}

func (e *Rxe) GetFromClassDefs(name string) (any, bool) {
	v, ok := e.classDefs[name]
	return v, ok
}

func (e *Rxe) GetFromPackageDefs(name string) (any, bool) {
	v, ok := e.packageDefs[name]
	return v, ok
}

func (e *Rxe) SetArgs(name string, val any) {
	e.callArgs[name] = val.(Node)
}

func (e *Rxe) DeleteArgs(names ...string) {
	for _, name := range names {
		delete(e.callArgs, name)
	}
}

func (e *Rxe) DumpRxe() RxeHandler {
	rxe := &Rxe{
		variables:     e.variables.Dump(),
		callArgs:      e.callArgs.Dump(),
		classDefs:     make(map[string]*ClassDef, 32),
		funcDefs:      make(map[string]*FunctionDef, 32),
		packageDefs:   make(map[string]*Package, 32),
		callingObject: e.callingObject,
		argsStack:     nil,
		goroutineChan: e.goroutineChan,
	}

	// copy classDefs
	for k, v := range e.classDefs {
		rxe.classDefs[k] = v
	}

	// copy funcDefs
	for k, v := range e.funcDefs {
		rxe.funcDefs[k] = v
	}

	// copy packageDefs
	for k, v := range e.packageDefs {
		rxe.packageDefs[k] = v
	}

	// copy argsStack
	for _, v := range e.argsStack {
		rxe.argsStack = append(rxe.argsStack, v.Dump())
	}

	// return
	return rxe
}
