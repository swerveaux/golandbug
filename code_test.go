package golandbug

import "testing"

type somethingDoer interface {
	doSomething()
}

func helperFunction(doer somethingDoer) {}

func TestSomeStruct(t *testing.T) {
	var something SomeStruct
	helperFunction(&something)
}
