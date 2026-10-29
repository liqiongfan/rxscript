//
// The RXScript author
//

package ast

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "embed"
)

const (
	OmitMore = `...`
	ThisVar  = `this`
)

// stdLibrary is the RXScript programming language's standard library, which written in RXScript.
var (
	//go:embed std/*
	stdLibrary embed.FS
)

// Lex token split method
// yacc will invoke Lex method to get next token
func (lex *RxeLexer) Lex(lval *rxeSymType) int {
	token := getToken(lex)
	//fmt.Printf("token: %d, %s, %v\n", token, rxeSymNames[rxeXLAT[token]], lex.value)
	switch token {
	case INT, STRING, BOOL, DOUBLE, OMITMORE, NAME, PATH, SYSTEMPATH:
		lval.val = lex.value
	case NULL:
		lval.val = nil
	}
	return token
}

// If yacc or re2go occur mistake, Error will be invoked.
func (lex *RxeLexer) Error(s string) {
	fmt.Printf("RXScript: %s in %s on line %d\n", s, lex.fileName, lex.line)
}

// RunScript RXEngine entrypoint method
// each script run start from here. debug mode give you more bison reduce/shift rule details
func RunScript(scriptName string, debug, goNum int) int {
	rxeDebug = debug

	if filepath.Ext(scriptName) != ScriptSuffix {
		scriptName += ScriptSuffix
	}

	buff, err := os.ReadFile(scriptName)
	if err != nil {
		CheckSystemError(&SystemErrorDef{Value: "RXScript: " + err.Error()}, nil)
	}

	// re2go lexer
	lex := &RxeLexer{
		buff:     string(buff) + "\x00", // file content end with \x00
		line:     1,                     // default line
		column:   1,                     // default column
		fileName: scriptName,            // file name
	}

	// yacc function
	result := rxeParse(lex)
	if result != 0 {
		return result
	}

	// RXEngine data structure
	engine := &Rxe{
		variables:     NewNodeStore(),
		callArgs:      NewNodeStore(),
		classDefs:     make(map[string]*ClassDef, 64),
		funcDefs:      make(map[string]*FunctionDef, 64),
		packageDefs:   make(map[string]*Package, 64),
		goroutineChan: make(chan struct{}, goNum),
	}

	// deal with import directives
	for _, stmt := range lex.result {
		switch v := stmt.(type) {
		case *Import:
			// each command after Execute, should use CheckSystemError to check system error
			CheckSystemError(v.Execute(engine), engine)
		}
	}

	// register global function and class into environment
	for _, stmt := range lex.result {
		switch v := stmt.(type) {
		case *ClassDef:
			engine.classDefs[v.Name] = v
		case *FunctionDef:
			if v.Name != `main` {
				engine.funcDefs[v.Name] = v
			}
		}
	}

	// Run `main` method
	for _, stmt := range lex.result {
		switch v := stmt.(type) {
		case *FunctionDef:
			if v.Name == `main` {
				switch ret := CheckSystemError(v.Invoke(nil, engine), engine).(type) {
				case int:
					return ret
				case *IntDef:
					return int(ret.Value)
				default:
					// any other things will return system 1, which means error occurs.
					return 1
				}
			}
		}
	}

	log.Fatal("RXScript: 'main' not defined.")
	return 0
}
