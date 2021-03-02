package leetcode

import (
	"testing"

	"github.com/rosenlo/leetcode/structure/tree"
)

func TestSameTree(t *testing.T) {
	tests := [][2][]int{
		{{1, 2, 3}, {1, 2, 3}},
		{{1, 2}, {1, 2}},
		{{1, 2}, {1, 3, 2}},
		{{1, 2, 1}, {1, 2, 1}},
		{{}, {}},
	}
	results := []bool{
		true,
		true,
		false,
		true,
		true,
	}
	for i := 0; i < len(tests); i++ {
		p := tree.Ints2TreeNode(tests[i][0])
		q := tree.Ints2TreeNode(tests[i][1])
		ret := isSameTree(p, q)
		t.Log(tests[i], ret, results[i])
		if ret != results[i] {
			t.Fatalf("Wrong Answer, ret :%v right ret: %v", ret, results[i])
		}
	}
}
