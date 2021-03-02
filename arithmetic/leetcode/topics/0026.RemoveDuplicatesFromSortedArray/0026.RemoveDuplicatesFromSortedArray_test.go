package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := [][]int{
		[]int{1, 1, 2},
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		[]int{},
	}
	results := []int{
		2,
		5,
		0,
	}
	for i := 0; i < len(tests); i++ {
		if ret := removeDuplicates(tests[i]); ret != results[i] {
			t.Fatalf("Wrong Answer, ret: %v, right ret: %v", ret, results[i])
		}
	}
}
