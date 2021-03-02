package leetcode

import (
	"reflect"
	"testing"

	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
)

func TestBalanced(t *testing.T) {
	tests := [][]int{
		{1, tree.NULL, 2, tree.NULL, 3},
		{3, 9, 20, tree.NULL, tree.NULL, 15, 7},
		{1, 2, 2, 3, 3, tree.NULL, tree.NULL, 4, 4},
		{1},
		{},
		{1, 2, 2, 3, tree.NULL, tree.NULL, 3, 4, tree.NULL, tree.NULL, 4},
	}
	results := []bool{
		false,
		true,
		false,
		true,
		true,
		false,
	}
	for i := 0; i < len(tests); i++ {
		ret := isBalanced(tree.Ints2TreeNode(tests[i]))
		if !reflect.DeepEqual(results[i], ret) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret, results[i])
		}
	}
}
