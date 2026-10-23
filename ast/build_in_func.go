//
// The RXScript author
// built-in functions

package ast

// fetchBuildInFunc functions which build-in rxt engine.
// all built-in functions bellow here
func fetchBuildInFunc(name string) any {
	switch name {
	case `println`:
		return NewPrintln()
	case `listDir`:
		return NewListDir()
	case `getTimeResource`, `time`:
		return NewTime()
	case `getLocation`:
		return NewLocation()
	case `isInt`:
		return NewIsInt()
	case `isFloat`, `isDouble`:
		return NewIsDouble()
	case `isNull`:
		return NewIsNull()
	case `isBool`, `isBoolean`:
		return NewIsBool()
	case `isString`:
		return NewIsString()
	case `isObject`:
		return NewIsObject()
	case `isArray`:
		return NewIsArray()
	case `isDictionary`, `isDict`:
		return NewIsDictionary()
	case `isFunction`, `isFunc`:
		return NewIsFunction()
	case `readFile`:
		return NewReadFile()
	case `removeFile`, `delFile`:
		return NewRemoveFile()
	case `putFile`:
		return NewPutFile()
	case `rename`:
		return NewRename()
	case `sleep`:
		return NewSleep()
	case `series`:
		return NewRange()
	case `channel`:
		return NewChannel()
	case `wait`:
		return NewWait()
	case `waitAdd`:
		return NewWaitAdd()
	case `waitDone`:
		return NewWaitDone()
	case `waitOk`:
		return NewWaitOk()
	case `atomicInt`:
		return NewAtomicInt()
	case `storeAtomicInt`:
		return NewAtomicIntStore()
	case `readAtomicInt`:
		return NewAtomicIntLoad()
	case `trim`:
		return NewTrim()
	case `trimLeft`:
		return NewTrimLeft()
	case `trimRight`:
		return NewTrimRight()
	case `index`:
		return NewIndex()
	case `lastIndex`:
		return NewLastIndex()
	case `format`:
		return NewFormat()
	case `explode`:
		return NewExplode()
	case `join`:
		return NewJoin()
	case `newHttp`:
		return NewHttpRequestFunc()
	case `setHttpHeader`:
		return NewSetHttpHeaderFunc()
	case `httpDo`:
		return NewHttpDoFunc()
	case `newTimer`:
		return NewNewTimer()
	case `resetTimer`:
		return NewResetTimer()
	case `stopTimer`:
		return NewStopTimer()
	case `stopTicker`:
		return NewStopTicker()
	case `resetTicker`:
		return NewResetTicker()
	case `newTicker`:
		return NewNewTicker()
	default:
		return nil
	}
}
