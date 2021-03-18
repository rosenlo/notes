package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rosenlo/toolkits/structure/linkedlist"
)

type param struct {
	one    []int
	two    int
	result []int
}

func TestMergeLinkedList(t *testing.T) {
	tests := []param{
		{
			[]int{1, 2, 3, 4, 5},
			1,
			[]int{5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			0,
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5},
			6,
			[]int{},
		},
	}
	for i := 0; i < len(tests); i++ {
		pHead1 := linkedlist.Ints2ListNode(tests[i].one)
		ret := FindKthToTail(pHead1, tests[i].two)
		retNode := linkedlist.ListNode2Ints(ret)
		if !reflect.DeepEqual(retNode, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v %v, actual: %v expected: %v", tests[i].one, tests[i].two, retNode, tests[i].result)
		}
		fmt.Printf("[input]: %v %v\t", tests[i].one, tests[i].two)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", retNode)
	}

}
