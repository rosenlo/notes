package bst

import (
	"fmt"
	"testing"

	"github.com/rosenlo/leetcode/structure/tree"
)

type param struct {
	one    []int
	result int
}

func TestGetMinimumDifference(t *testing.T) {
	tests := []param{
		{
			[]int{1, 2, 3, 4, 5},
			3,
		},
		{
			[]int{1, 2, tree.NULL, 3, 4},
			2,
		},
		{
			[]int{4, -7, -3, tree.NULL, tree.NULL, -9, -3, 9, -7, -4, tree.NULL, 6, tree.NULL, -6, -6, tree.NULL, tree.NULL, 0, 6, 5, tree.NULL, 9, tree.NULL, tree.NULL, -1, -4, tree.NULL, tree.NULL, tree.NULL, -2},
			8,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := diameterOfBinaryTree2(tree.Ints2TreeNode(tests[i].one))
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
