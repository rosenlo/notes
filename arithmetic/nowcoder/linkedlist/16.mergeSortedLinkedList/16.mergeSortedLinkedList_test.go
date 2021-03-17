package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rosenlo/toolkits/structure/linkedlist"
)

type param struct {
	one    []int
	two    []int
	result []int
}

func TestMergeLinkedList(t *testing.T) {
	tests := []param{
		{
			[]int{1, 3, 4},
			[]int{2, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{2, 3, 4},
			[]int{1, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{2, 4, 6},
			[]int{1, 3},
			[]int{1, 2, 3, 4, 6},
		},
	}
	for i := 0; i < len(tests); i++ {
		pHead1 := linkedlist.Ints2ListNode(tests[i].one)
		pHead2 := linkedlist.Ints2ListNode(tests[i].two)
		ret := Merge(pHead1, pHead2)
		mergeHead := linkedlist.ListNode2Ints(ret)
		if !reflect.DeepEqual(mergeHead, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v %v, actual: %v expected: %v", tests[i].one, tests[i].two, mergeHead, tests[i].result)
		}
		fmt.Printf("[input]: %v %v\t", tests[i].one, tests[i].two)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", mergeHead)
	}

}
