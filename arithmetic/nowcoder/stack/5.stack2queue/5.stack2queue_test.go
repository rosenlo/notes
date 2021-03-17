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

func TestStack2Queue(t *testing.T) {
	tests := []param{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 1, 2, 3, 4},
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := []int{}

		for x := 0; x < len(tests[i].one); x++ {
			fmt.Println("push", tests[i].one[x])
			Push(tests[i].one[x])
			if x%2 == 0 {
				node := Pop()
				fmt.Println("pop", node)
				ret = append(ret, node)
			}
		}
		if !reflect.DeepEqual(ret, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
