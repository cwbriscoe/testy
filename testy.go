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
	v1 = normalizeValue(v1)
	v2 = normalizeValue(v2)
	if v1 != v2 {
		tb.Fatalf("%s:%d NOTEQUAL: expected %v to equal %v", file, line, v1, v2)
	}
}

// NotEquals fails if v1 equals v2
func NotEquals(tb testing.TB, v1, v2 interface{}) {
	_, file, line, _ := runtime.Caller(1)
	v1 = normalizeValue(v1)
	v2 = normalizeValue(v2)
	if v1 == v2 {
		tb.Fatalf("%s:%d EQUAL: expected %v to not equal %v", file, line, v1, v2)
	}
}

// Nil fails if v1 is not nil
func Nil(tb testing.TB, v1 interface{}) {
	_, file, line, _ := runtime.Caller(1)
	v1 = normalizeValue(v1)
	if v1 != nil {
		tb.Fatalf("%s:%d NOTNIL: expected %v to be <nil>", file, line, v1)
	}
}

// NotNil fails if v1 is nil
func NotNil(tb testing.TB, v1 interface{}) {
	_, file, line, _ := runtime.Caller(1)
	v1 = normalizeValue(v1)
	if v1 == nil {
		tb.Fatalf("%s:%d NOTNIL: expected %v to be not be <nil>", file, line, v1)
	}
}

// Assert fails if the condition expression is false.
func Assert(tb testing.TB, condition bool, msg string) {
	_, file, line, _ := runtime.Caller(1)
	if !condition {
		tb.Fatalf("%s:%d ASSERT: %s", file, line, msg)
	}
}

// normalizeValue converts []byte to nil for comparison
func normalizeValue(v1 interface{}) interface{} {
	switch v1.(type) {
	case int64:
		return v1
	case int32:
		return int64(v1.(int32))
	case int16:
		return int64(v1.(int16))
	case int8:
		return int64(v1.(int8))
	case uint64:
		return v1
	case uint32:
		return uint64(v1.(uint32))
	case uint16:
		return uint64(v1.(uint16))
	case uint8:
		return uint64(v1.(uint8))
	case float64:
		return v1
	case float32:
		return float64(v1.(float32))
	case complex128:
		return v1
	case complex64:
		return complex128(v1.(complex64))
	case []byte:
		if len(v1.([]byte)) == 0 {
			return nil
		}
	default:
		break
	}
	return v1
}
