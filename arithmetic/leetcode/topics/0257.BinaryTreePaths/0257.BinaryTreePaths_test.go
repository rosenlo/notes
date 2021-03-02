package leetcode

import (
	"reflect"
	"testing"

	"github.com/rosenlo/leetcode/structure/tree"
)

type param struct {
	one    []int
	result []string
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []param{
		{
			[]int{},
			[]string{},
		},
		{
			[]int{1, 2, 3, tree.NULL, 5},
			[]string{"1->2->5", "1->3"},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"},
		},
		{
			[]int{6, 1, tree.NULL, tree.NULL, 3, 2, 5, tree.NULL, tree.NULL, 4},
			[]string{"6->1->3->2", "6->1->3->5->4"},
		},
	}
	for i := 0; i < len(tests); i++ {
		root := tree.Ints2TreeNode(tests[i].one)
		ret := binaryTreePaths(root)
		if !reflect.DeepEqual(ret, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		ret2 := binaryTreePaths2(root)
		if !reflect.DeepEqual(ret2, tests[i].result) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret2, tests[i].result)
		}
	}
}
