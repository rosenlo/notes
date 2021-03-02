package leetcode

import (
	"testing"

	"github.com/rosenlo/leetcode/structure/tree"
)

func TestMinDepth(t *testing.T) {
	tests := [][]int{
		{3, 9, 20, tree.NULL, tree.NULL, 15, 7},
		{2, tree.NULL, 3, tree.NULL, 4, tree.NULL, 5, tree.NULL, 6},
		{1, 2},
		{1},
		{1, tree.NULL, 2},
		{-9, -3, 2, tree.NULL, 4, 4, 0, -6, tree.NULL, -5},
	}
	results := []int{
		2,
		5,
		2,
		1,
		2,
		3,
	}
	for i := 0; i < len(tests); i++ {
		root := tree.Ints2TreeNode(tests[i])
		ret := minDepth(root)
		if ret != results[i] {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret, results[i])
		}
	}
}
