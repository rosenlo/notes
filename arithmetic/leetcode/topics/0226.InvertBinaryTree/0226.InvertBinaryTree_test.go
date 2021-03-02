package leetcode

import (
	"reflect"
	"testing"

	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
)

type param struct {
	one    []int
	result []int
}

func TestInvertTree(t *testing.T) {
	tests := []param{
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			[]int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			[]int{4, 2, 7, 1, 3},
			[]int{4, 7, 2, tree.NULL, tree.NULL, 3, 1},
		},
		{
			[]int{4, 2, 7, tree.NULL, tree.NULL, 6, 9},
			[]int{4, 7, 2, 9, 6},
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := invertTree(tree.Ints2TreeNode(tests[i].one))
		if !reflect.DeepEqual(tree.TreeNode2Ints(ret), tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
	}
}
