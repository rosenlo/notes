package main

import (
	"fmt"
	"reflect"
	"testing"
)

type param struct {
	one    []int
	result []int
}

func TestMultiply(t *testing.T) {
	tests := []param{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{120, 60, 40, 30, 24},
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := multiply(tests[i].one)
		if !reflect.DeepEqual(ret, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
