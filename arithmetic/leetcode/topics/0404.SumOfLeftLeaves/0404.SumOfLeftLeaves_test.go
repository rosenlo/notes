package leetcode

import (
	"testing"

	"github.com/rosenlo/leetcode/structure/tree"
)

type param struct {
	one    []int
	result int
}

func TestSumOfLeftLeaves(t *testing.T) {
	tests := []param{
		{
			[]int{3, 9, 20, tree.NULL, tree.NULL, 15, 7},
			24,
		},
		{
			[]int{3, 9, 20, 11, tree.NULL, 15, 7},
			26,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := sumOfLeftLeaves(tree.Ints2TreeNode(tests[i].one))
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
	}
}
