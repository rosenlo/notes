package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rosenlo/toolkits/structure/linkedlist"
)

type param struct {
	one    []int
	result []int
}

func TestMergeLinkedList(t *testing.T) {
	tests := []param{
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
	}
	for i := 0; i < len(tests); i++ {
		pHead1 := linkedlist.Ints2ListNode(tests[i].one)
		ret := ReverseList(pHead1)
		retNode := linkedlist.ListNode2Ints(ret)
		if !reflect.DeepEqual(retNode, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, retNode, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", retNode)
	}

}
