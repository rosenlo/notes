package bst

import (
	"fmt"
	"testing"

	"github.com/rosenlo/notes/arithmetic/leetcode/structure/tree"
)

type param struct {
	one    []int
	result int
}

func TestGetMinimumDifference(t *testing.T) {
	tests := []param{
		{
			[]int{1, tree.NULL, 3, 2, 4, tree.NULL, 5, 6},
			3,
		},
		{
			[]int{1, tree.NULL, 2, 3, 4, 5, tree.NULL, tree.NULL, 6, 7, tree.NULL, 8, tree.NULL, 9, 10, tree.NULL, tree.NULL, 11, tree.NULL, 12, tree.NULL, 13, tree.NULL, tree.NULL, 14},
			5,
		},
		{
			[]int{44},
			1,
		},
		{
			[]int{},
			0,
		},
	}
	for i := 0; i < len(tests); i++ {
		root := tree.Ints2NaryTreeNode(tests[i].one)
		ret := maxDepth(root)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
