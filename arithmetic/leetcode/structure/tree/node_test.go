package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInts2TreeNode(t *testing.T) {
	tests := [][]int{
		{1, 2, 3},
		{1, 1, 2},
		{1, 2, 3, 4, 5},
		{},
		{3, 9, 20, NULL, NULL, 15, 7},
		{
			3,
			4, 5,
			-7, -6, NULL, NULL,
			-7, NULL, -5, NULL, NULL, NULL, -4,
		},
	}
	for i := 0; i < len(tests); i++ {
		node := Ints2TreeNode(tests[i])
		ret := TreeNode2Ints(node)
		t.Log(tests[i], ret)
		if !reflect.DeepEqual(tests[i], ret) {
			t.Fatalf("Wrong Answer, ret: %v right ret: %v", ret, tests[i])
		}
	}
}

func TestTreeOrder(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6},
	}
	for i := 0; i < len(tests); i++ {
		node := Ints2TreeNode(tests[i])
		fmt.Printf("PreOrder:  %v\n", PreOrder(node))
		fmt.Printf("InOrder:   %v\n", InOrder(node))
		fmt.Printf("PostOrder: %v\n", PostOrder(node))
	}
}
