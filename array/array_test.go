package array

import (
	"testing"
	"fmt"
	"reflect"
)

const length = 3
const capacity = 3

func TestArray0(t *testing.T) {
	//// Nil slice has length and capacity of 0
	//var a []int
	//
	//// Empty array has length of 2 and capacity of 2
	//var d [2]int

	x0 := Array0()

	assertEqual(t, "[3]string", reflect.TypeOf(x0).String())
	assertEqual(t, capacity, cap(x0))
	assertEqual(t, length, len(x0))
}

func TestArray1(t *testing.T) {
	x1 := Array1()

	assertEqual(t, "[3]string", reflect.TypeOf(x1).String())
	assertEqual(t, capacity, cap(x1))
	assertEqual(t, length, len(x1))
}

func TestArray2(t *testing.T) {
	x2 := Array2()

	assertEqual(t, "[3]string", reflect.TypeOf(x2).String())
	assertEqual(t, capacity, cap(x2))
	assertEqual(t, length, len(x2))
}

func TestArray3(t *testing.T) {
	x3 := Array3()

	assertEqual(t, "[3]string", reflect.TypeOf(x3).String())
	assertEqual(t, capacity, cap(x3))
	assertEqual(t, length, len(x3))
}


func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	message := fmt.Sprintf("%v != %v", a, b)
	t.Fatal(message)
}

