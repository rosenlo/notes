package main

import (
	"testing"

	"github.com/rosenlo/toolkits/structure/tree"
)

type param struct {
	one    []int
	result int
}

func TestTreeDepth(t *testing.T) {
	tests := []param{
		{
			[]int{1, 2, 3, 4, 5, tree.NULL, 6, tree.NULL, tree.NULL, 7},
			4,
		},
	}
	for i := 0; i < len(tests); i++ {
		root := tree.Ints2TreeNode(tests[i].one)
		ret := TreeDepth(root)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, ret: %v right ret: %v", ret, tests[i].result)
		}
	}
}
