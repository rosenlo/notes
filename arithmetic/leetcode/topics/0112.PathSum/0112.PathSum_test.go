package leetcode

import (
	"testing"

	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
)

type param struct {
	one    []int
	two    int
	result bool
}

func TestMinDepth(t *testing.T) {
	tests := []param{
		{
			[]int{5, 4, 8, 11, tree.NULL, 13, 4, 7, 2, tree.NULL, tree.NULL, tree.NULL, 1},
			22,
			true,
		},
		{
			[]int{1, 2, 3},
			5,
			false,
		},
		{
			[]int{1, 2},
			0,
			false,
		},
		{
			[]int{1, 2},
			1,
			false,
		},
		{
			[]int{1},
			1,
			true,
		},
		{
			[]int{1},
			0,
			false,
		},
		{
			[]int{1},
			0,
			false,
		},
		{
			[]int{0},
			0,
			true,
		},
		{
			[]int{},
			0,
			false,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := hasPathSum(tree.Ints2TreeNode(tests[i].one), tests[i].two)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
	}
}
