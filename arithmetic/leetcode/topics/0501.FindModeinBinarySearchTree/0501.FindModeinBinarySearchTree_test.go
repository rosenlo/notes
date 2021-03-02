package leetcode

import (
	"fmt"
	"testing"

	"github.com/rosenlo/notes/arithmetic/leetcode/structure/tree"
)

type param struct {
	one    []int
	result []int
}

func TestFindMode(t *testing.T) {
	tests := []param{
		{
			[]int{1, tree.NULL, 2, 2},
			[]int{2},
		},
		{
			[]int{1, tree.NULL, 2, 2, 3, tree.NULL, tree.NULL, 3},
			[]int{2, 3},
		},
		{
			[]int{2, 1, 3, tree.NULL, tree.NULL, tree.NULL, 4},
			[]int{4, 2, 1, 3},
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, tree.NULL, tree.NULL, 2, 6},
			[]int{2, 6},
		},
		{
			[]int{-2, -2, -2},
			[]int{-2},
		},
		{
			[]int{2, 1, tree.NULL, tree.NULL, 2, 1, tree.NULL, tree.NULL, 2},
			[]int{2},
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := findMode(tree.Ints2TreeNode(tests[i].one))
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}
}
