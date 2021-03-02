package leetcode

import (
	"testing"

	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
)

type param struct {
	one    []int
	two    []int
	three  []int
	result int
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []param{
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{2},
			[]int{8},
			6,
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{2},
			[]int{4},
			2,
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{7},
			[]int{9},
			8,
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{2},
			[]int{3},
			2,
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{3},
			[]int{8},
			6,
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{0},
			[]int{8},
			6,
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{0},
			[]int{3},
			2,
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 3, 5},
			[]int{3},
			[]int{8},
			6,
		},
		{
			[]int{2, 1},
			[]int{2},
			[]int{1},
			2,
		},
		{
			[]int{2, tree.NULL, 1},
			[]int{2},
			[]int{1},
			2,
		},
		{
			[]int{3, 1, 4, tree.NULL, 2},
			[]int{2},
			[]int{3},
			3,
		},
		{
			[]int{3, 1, 4, tree.NULL, 2, tree.NULL, 5},
			[]int{1},
			[]int{5},
			3,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := lowestCommonAncestor(tree.Ints2TreeNode(tests[i].one), tree.Ints2TreeNode(tests[i].two), tree.Ints2TreeNode(tests[i].three))
		if ret.Val != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret.Val, tests[i].result)
		}
	}
}
