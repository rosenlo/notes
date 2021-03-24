package main

import (
	"fmt"
	"reflect"
	"testing"
)

type param struct {
	one    []int
	two    []int
	result bool
}

func TestIsPopOrder(t *testing.T) {
	tests := []param{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 1, 2},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{4, 6, 5, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 2, 6, 5},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 4, 6, 5, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 6, 4, 5, 2, 1},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 4, 3, 5, 2},
			true,
		},
	}
	for i := 0; i < len(tests); i++ {

		ret := IsPopOrder(tests[i].one, tests[i].two)

		if !reflect.DeepEqual(ret, tests[i].result) {
			t.Fatalf("Wrong Answer, [intpu]: %v %v, actual: %v expected: %v", tests[i].one, tests[i].two, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v %v\t", tests[i].one, tests[i].two)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
