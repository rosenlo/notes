package main

import (
	"fmt"
	"testing"
)

type param struct {
	one    []int
	result int
}

func TestFibonacci(t *testing.T) {
	tests := []param{
		{
			[]int{3, 4, 5, 1, 2},
			1,
		},
		{
			[]int{5, 8, 9, 2, 3},
			2,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := minNumberInRotateArray(tests[i].one)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
