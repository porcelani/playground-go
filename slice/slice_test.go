package slice

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {

	x1  := Slice1(); fmt.Println(x1); fmt.Println(len(x1)); fmt.Println(cap(x1))
}


