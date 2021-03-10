package main

import (
	"reflect"
	"testing"

	"github.com/rosenlo/toolkits/structure/tree"
)

type param struct {
	one    []int
	result []int
}

func TestSymeetricTree(t *testing.T) {
	tests := []param{
		{
			[]int{8, 6, 10, 5, 7, 9, 11},
			[]int{8, 10, 6, 11, 9, 7, 5},
		},
	}
	for i := 0; i < len(tests); i++ {
		root := tree.Ints2TreeNode(tests[i].one)
		ret := tree.TreeNode2Ints(Mirror(root))
		if !reflect.DeepEqual(ret, tests[i].result) {
			t.Fatalf("Wrong Answer, ret: %v right ret: %v", ret, tests[i].result)
		}
	}
}
