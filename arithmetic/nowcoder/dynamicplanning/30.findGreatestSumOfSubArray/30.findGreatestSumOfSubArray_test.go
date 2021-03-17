package main

import (
	"fmt"
	"testing"
)

type param struct {
	one    []int
	result int
}

func TestFindGreatestSumOfSubArray(t *testing.T) {
	tests := []param{
		{
			[]int{1, -2, 3, 10, -4, 7, 2, -5},
			18,
		},
		{
			[]int{-2, -1},
			-1,
		},
		{
			[]int{-1},
			-1,
		},
		{
			[]int{2, -3, 1, 3, -3, 2, 2, 1},
			6,
		},
		{
			[]int{1, -2, 3, 10, -4, 7, 2, -5},
			18,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := FindGreatestSumOfSubArray2(tests[i].one)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
