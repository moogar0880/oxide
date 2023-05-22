package assert

import (
	"bytes"
	"reflect"
	"testing"
)

type TestingT interface {
	*testing.T | *testing.B | *mockTest

	Errorf(format string, args ...interface{})
}

type mockTest struct{}

func (t *mockTest) Errorf(_ string, _ ...interface{}) {}

func Equal[T TestingT](t T, expected, actual interface{}) bool {
	if isFunction(expected) || isFunction(actual) {
		t.Errorf("you can not compare functions for equality")
		return false
	}

	if !equal(expected, actual) {
		t.Errorf("actual value did not match expected\nexpected: %v\n actual=%v", expected, actual)
		return false
	}

	return true
}

func equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}

	return bytes.Equal(exp, act)
}

func isFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}
