package slice

import (
	"testing"
	"fmt"
	"reflect"
)

const expect = 3
const capacity = 3

func TestSlice0(t *testing.T) {
	country, code := Slice0()

	assertEqual(t, "[]string", reflect.TypeOf(country).String())
	assertEqual(t, 0, cap(country))
	assertEqual(t, 0, len(country))

	assertEqual(t, "[]int", reflect.TypeOf(code).String())
	assertEqual(t, 0, cap(code))
	assertEqual(t, 0, len(code))

}

func TestSlice1(t *testing.T) {
	x1 := Slice1()

	assertEqual(t, "[]string", reflect.TypeOf(x1).String())
	assertEqual(t, capacity, cap(x1))
	assertEqual(t, expect, len(x1))
}

func TestSliceMake(t *testing.T) {
	sliceMaked := SliceMake()

	assertEqual(t, "[]string", reflect.TypeOf(sliceMaked).String())
	assertEqual(t, 5, cap(sliceMaked))
	assertEqual(t, 5, len(sliceMaked))

	expect := 0;
	for country := range sliceMaked {
		assertEqual(t, expect, country)
		assertEqual(t, "", sliceMaked[expect])
		expect++
	}
}

func TestSliceMakeWithAdd(t *testing.T) {
	sliceMaked := SliceMakeWithAdd()

	assertEqual(t, "[]string", reflect.TypeOf(sliceMaked).String())
	assertEqual(t, 5, cap(sliceMaked))
	assertEqual(t, 5, len(sliceMaked))

	assertEqual(t, "Brazil", sliceMaked[0])
	assertEqual(t, "USA", sliceMaked[1])
	assertEqual(t, "", sliceMaked[2])
	assertEqual(t, "", sliceMaked[3])
	assertEqual(t, "", sliceMaked[4])
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	message := fmt.Sprintf("%v != %v", a, b)
	t.Fatal(message)
}

