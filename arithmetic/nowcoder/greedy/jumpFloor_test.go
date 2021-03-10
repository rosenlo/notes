package main

import (
	"fmt"
	"testing"
)

type param struct {
	one    int
	result int
}

func TestJumpFloorII(t *testing.T) {
	tests := []param{
		{
			1,
			1,
		},
		{
			3,
			4,
		},
		{
			4,
			8,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := jumpFloorII(tests[i].one)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
