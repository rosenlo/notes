package leetcode

import (
	"reflect"
	"testing"

	"github.com/rosenlo/leetcode/structure/tree"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := [][]int{
		{-10, -3, 0, 5, 9},
		{-9, -6, 0, 3, 7, 11, 13},
		{0, 1, 2, 3, 4, 5},
		{1},
		{},
	}
	results := [][]int{
		{0, -3, 9, -10, tree.NULL, 5},
		{3, -6, 11, -9, 0, 7, 13},
		{3, 1, 5, 0, 2, 4},
		{1},
		{},
	}
	for i := 0; i < len(tests); i++ {
		ret := tree.TreeNode2Ints(sortedArrayToBST(tests[i]))
		if !reflect.DeepEqual(results[i], ret) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret, results[i])
		}
	}
}
