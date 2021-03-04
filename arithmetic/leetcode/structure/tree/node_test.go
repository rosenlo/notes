package tree

import (
	"fmt"
	"reflect"
	"testing"
)

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
		fmt.Printf("[input]: %v\t", tests[i])
		fmt.Printf("[expect]: %v\t", tests[i])
		fmt.Printf("[output]: %v\t\n", ret)
		if !reflect.DeepEqual(tests[i], ret) {
			t.Fatalf("Wrong Answer, ret: %v right ret: %v", ret, tests[i])
		}
	}
}

func TestInts2NaryTreeNode(t *testing.T) {
	tests := [][]int{
		{1, NULL, 3, 2, 4, NULL, 5, 6},
		{1, NULL, 2, 3, 4, 5, NULL, NULL, 6, 7, NULL, 8, NULL, 9, 10, NULL, NULL, 11, NULL, 12, NULL, 13, NULL, NULL, 14},
	}
	for i := 0; i < len(tests); i++ {
		node := Ints2NaryTreeNode(tests[i])
		ret := NaryTreeNode2Ints(node)
		fmt.Printf("[input]: %v\t", tests[i])
		fmt.Printf("[expect]: %v\t", tests[i])
		fmt.Printf("[output]: %v\t\n", ret)
		if !reflect.DeepEqual(tests[i], ret) {
			t.Fatalf("Wrong Answer, actual: %v expected: %v", ret, tests[i])
		}
	}
}
