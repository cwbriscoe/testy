// Copyright 2020 Christopher Briscoe.  All rights reserved.

package testy

import (
	"errors"
	"testing"
)

func TestOk(t *testing.T) {
	var err error
	Ok(t, err)
}

func TestNotOk(t *testing.T) {
	err := errors.New("test error")
	NotOk(t, err)
}

func TestEquals(t *testing.T) {
	Equals(t, 1, 1)
	Equals(t, 0, 0)
	Equals(t, 9.9, 9.9)
	Equals(t, "a", "a")
	Equals(t, "A", "A")
	Equals(t, "Foo", "Foo")

	v1 := 1
	v2 := 0
	v3 := 9.9
	v4 := "a"
	v5 := "A"
	v6 := "Foo"

	Equals(t, 1, v1)
	Equals(t, 0, v2)
	Equals(t, 9.9, v3)
	Equals(t, "a", v4)
	Equals(t, "A", v5)
	Equals(t, "Foo", v6)
}

func TestNotEquals(t *testing.T) {
	NotEquals(t, 1, 2)
	NotEquals(t, 0, 1)
	NotEquals(t, 9.9, 9.8)
	NotEquals(t, "a", "b")
	NotEquals(t, "A", "B")
	NotEquals(t, "Foo", "Bar")

	v1 := 1
	v2 := 0
	v3 := 9.9
	v4 := "a"
	v5 := "A"
	v6 := "Foo"

	NotEquals(t, 2, v1)
	NotEquals(t, 1, v2)
	NotEquals(t, 9.8, v3)
	NotEquals(t, "b", v4)
	NotEquals(t, "B", v5)
	NotEquals(t, "Bar", v6)
}

func TestNil(t *testing.T) {
	Nil(t, nil)
}

func TestNotNil(t *testing.T) {
	v := 1
	NotNil(t, v)
}

func TestAssert(t *testing.T) {
	Assert(t, 1 == 1, "1 != 1")
	Assert(t, 0 == 0, "0 != 0")
	Assert(t, 9.9 == 9.9, "9.9 != 9.9")
	Assert(t, "a" == "a", "a != a")
	Assert(t, "A" == "A", "A != A")
	Assert(t, "Foo" == "Foo", "Foo != Foo")
	Assert(t, "Foo" == "Foo", "Foo != Foo")
}

func TestNilValues(t *testing.T) {
	var b1 []byte
	var b2 [10]byte
	b2[0] = 1
	Assert(t, b1 == nil, "b1 should be nil")
	NotEquals(t, nil, b2)
	Equals(t, nil, b1)
}

func TestNumericEqualities(t *testing.T) {
	var i1 int8 = 1
	var i2 int16 = 1
	var i3 int32 = 1
	var i4 int64 = 1
	var i5 rune = 1
	Equals(t, i1, i2)
	Equals(t, i1, i3)
	Equals(t, i1, i4)
	Equals(t, i1, i5)
	Equals(t, i2, i3)
	Equals(t, i2, i4)
	Equals(t, i2, i5)
	Equals(t, i3, i4)
	Equals(t, i3, i5)
	Equals(t, i4, i5)

	var u1 uint8 = 1
	var u2 uint16 = 1
	var u3 uint32 = 1
	var u4 uint64 = 1
	var u5 byte = 1
	Equals(t, u1, u2)
	Equals(t, u1, u3)
	Equals(t, u1, u4)
	Equals(t, u1, u5)
	Equals(t, u2, u3)
	Equals(t, u2, u4)
	Equals(t, u2, u5)
	Equals(t, u3, u4)
	Equals(t, u3, u5)
	Equals(t, u4, u5)

	var f1 float32 = 1.0
	var f2 float64 = 1.0
	Equals(t, f1, f2)

	var c1 complex64 = 2 + 3i
	var c2 complex128 = 2 + 3i
	Equals(t, c1, c2)
}
