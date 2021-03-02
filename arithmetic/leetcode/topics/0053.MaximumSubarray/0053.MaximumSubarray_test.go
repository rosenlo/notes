package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1, -7, 3, 9, -8, -2, 9, -4, 5, 9},
		{-4, 3, -1, 2, -5, -2},
		{-4, 3, -1, -1, -5, -2},
		{-1},
		{8, -19, 5, -4, 20},
		{2, -3, 1, 3, -3, 2, 2, 1},
		{-2, -1},
	}
	results := []int{
		6,
		21,
		4,
		3,
		-1,
		21,
		6,
		-1,
	}
	for i := 0; i < len(tests); i++ {
		if ret := maxSubArray2(tests[i]); ret != results[i] {
			t.Fatalf("Wrong Answer, ret: %v right ret: %v", ret, results[i])
		}
	}
}
