package main

import (
	"fmt"
	"testing"
)

type param struct {
	one    int
	two    int
	result int
}

func TestAdd(t *testing.T) {
	tests := []param{
		{
			1,
			1,
			2,
		},
		{
			3,
			4,
			7,
		},
		{
			1,
			9,
			10,
		},
		{
			0,
			1,
			1,
		},
		{
			100,
			900,
			1000,
		},
		{
			111,
			899,
			1010,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := Add(tests[i].one, tests[i].two)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, %v, actual: %v expected: %v", tests[i].one, tests[i].two, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v %v\t", tests[i].one, tests[i].two)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
