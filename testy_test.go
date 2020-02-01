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
