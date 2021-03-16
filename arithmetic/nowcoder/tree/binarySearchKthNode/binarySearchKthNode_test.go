package main

import (
	"reflect"
	"testing"

	"github.com/rosenlo/toolkits/structure/tree"
)

type param struct {
	one    []int
	two    int
	result []int
}

func TestKthNode(t *testing.T) {
	tests := []param{
		{
			[]int{5, 3, 7, 2, 4, 6, 8},
			3,
			[]int{4},
		},
		{
			[]int{5, 3, 7, 2, 4, 6, 8},
			7,
			[]int{8},
		},
		{
			[]int{5},
			1,
			[]int{5},
		},
		{
			[]int{},
			0,
			[]int{},
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := KthNode(tree.Ints2TreeNode(tests[i].one), tests[i].two)
		retInts := tree.TreeNode2Ints(ret)
		if !reflect.DeepEqual(retInts, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], retInts, tests[i].result)
		}
	}
}
