package leetcode

import (
	"testing"

	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
)

func TestMaxDepth(t *testing.T) {
	tests := [][]int{
		{3, 9, 20, tree.NULL, tree.NULL, 15, 7},
		{1, tree.NULL, 2},
		{1},
		{},
		{
			3,
			4, 5,
			-7, -6, tree.NULL, tree.NULL,
			-7, tree.NULL, -5, tree.NULL, tree.NULL, tree.NULL, -4,
		},
		{
			1,
			2, 2,
			3, 3, tree.NULL, tree.NULL,
			tree.NULL, 4,
			tree.NULL, tree.NULL, 5, 5,
		},
	}
	results := []int{
		3,
		2,
		1,
		0,
		5,
		5,
	}
	for i := 0; i < len(tests); i++ {
		root := tree.Ints2TreeNode(tests[i])
		ret := maxDepth(root)
		if ret != results[i] {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret, results[i])
		}
	}
}
