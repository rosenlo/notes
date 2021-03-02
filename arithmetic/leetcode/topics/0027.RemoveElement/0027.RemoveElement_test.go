package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := [][]int{
		[]int{3, 2, 2, 3},
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
	}
	vals := []int{
		3,
		2,
	}
	results := []int{
		2,
		5,
	}
	for i := 0; i < len(tests); i++ {
		if ret := removeElement(tests[i], vals[i]); ret != results[i] {
			t.Fatalf("Wrong Answer, ret: %v, right ret: %v", ret, results[i])
		}
	}
}
