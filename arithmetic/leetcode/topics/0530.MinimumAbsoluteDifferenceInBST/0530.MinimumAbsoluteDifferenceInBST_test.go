package leetcode

import (
	"fmt"
	"testing"

	"github.com/rosenlo/leetcode/structure/tree"
)

type param struct {
	one    []int
	result int
}

func TestGetMinimumDifference(t *testing.T) {
	tests := []param{
		{
			[]int{1, tree.NULL, 4, 3, 5},
			1,
		},
		{
			[]int{1, tree.NULL, 3, 2},
			1,
		},
		{
			[]int{1, tree.NULL, 3, 3, 4},
			0,
		},
		{
			[]int{1, 0},
			1,
		},
		{
			[]int{1, tree.NULL, 3},
			2,
		},
		{
			[]int{236, 104, 701, tree.NULL, 227, tree.NULL, 911},
			9,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := getMinimumDifference(tree.Ints2TreeNode(tests[i].one))
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
