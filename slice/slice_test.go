package slice

import (
	"testing"
	"fmt"
	"reflect"
)

const expect = 3
const capacity = 3

func TestSlice0(t *testing.T) {
	x0 := Slice0()

	assertEqual(t, "[]string", reflect.TypeOf(x0).String())
	assertEqual(t, 0, cap(x0))
	assertEqual(t, 0, len(x0))
}

func TestSlice1(t *testing.T) {
	x1 := Slice1()

	assertEqual(t, "[]string", reflect.TypeOf(x1).String())
	assertEqual(t, capacity, cap(x1))
	assertEqual(t, expect, len(x1))
}


func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	message := fmt.Sprintf("%v != %v", a, b)
	t.Fatal(message)
}

