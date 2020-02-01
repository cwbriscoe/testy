// Copyright 2020 Christopher Briscoe.  All rights reserved.

// Package testy consists of a few simple testing helper functions.
// You can shorten your typing by declaring the following variables
// at the top of your _test.go files:
// var ok = testy.Ok
// var notOk = testy.NotOk
// var equals = testy.Equals
// var notEquals = testy.NotEquals
// var assert = testy.Assert
package testy

import (
	"runtime"
	"testing"
)

// Ok fails if err is not nil
func Ok(tb testing.TB, err error) {
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		tb.Fatalf("%s:%d NOTOK: %v", file, line, err)
	}
}

// NotOk fails if err is nil
func NotOk(tb testing.TB, err error) {
	_, file, line, _ := runtime.Caller(1)
	if err == nil {
		tb.Fatalf("%s:%d OK: %v", file, line, err)
	}
}

// Equals fails if v1 does not equal v2
func Equals(tb testing.TB, v1, v2 interface{}) {
	_, file, line, _ := runtime.Caller(1)
	if v1 != v2 {
		tb.Fatalf("%s:%d NOTEQUAL: expected %v, got %v", file, line, v1, v2)
	}
}

// NotEquals fails if v1 equals v2
func NotEquals(tb testing.TB, v1, v2 interface{}) {
	_, file, line, _ := runtime.Caller(1)
	if v1 == v2 {
		tb.Fatalf("%s:%d EQUAL: did not expect %v to equal %v", file, line, v1, v2)
	}
}

// Assert fails if the condition expression is false.
func Assert(tb testing.TB, condition bool, msg string) {
	_, file, line, _ := runtime.Caller(1)
	if !condition {
		tb.Fatalf("%s:%d ASSERT: %s", file, line, msg)
	}
}
