package leetcode

import (
	"testing"

	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
)

func TestSymeetricTree(t *testing.T) {
	tests := [][]int{
		{1, 2, 2, 3, 4, 4, 3},
		{1, 2, 2, 4, 3, 4, 3},
		{1},
		{1, 2, 2},
		{1, 2, 2, 4},
		{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 7, 7, 7, 6, 5},
		{},
	}
	results := []bool{
		true,
		false,
		true,
		true,
		false,
		true,
		true,
	}
	for i := 0; i < len(tests); i++ {
		root := tree.Ints2TreeNode(tests[i])
		ret := isSymmetric(root)
		t.Log(tests[i], ret, results[i])
		if ret != results[i] {
			t.Fatalf("Wrong Answer, ret: %v right ret: %v", ret, results[i])
		}
	}
}
