package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testsNums1 := [][]int{
		{1, 2, 3, 0, 0, 0},
		{3, 4, 5, 6, 0, 0},
		{2, 4, 5, 6, 0, 0},
		{1},
		{0, 1, 3, 5, 0, 0},
		{0, 1, 0, 0, 0},
		{-111, -98, -1, 5, 0, 0, 0},
		{2, 0},
		{4, 5, 6, 0, 0, 0},
		{1, 0, 0},
	}
	testsNums2 := [][]int{
		{2, 5, 6},
		{2, 3},
		{2, 3},
		{},
		{1, 2},
		{0, 1, 2},
		{-3, -2, 6},
		{1},
		{1, 2, 3},
		{1, 2},
	}
	testsM := []int{
		3,
		4,
		4,
		1,
		4,
		2,
		4,
		1,
		3,
		1,
	}
	testsN := []int{
		3,
		2,
		2,
		0,
		2,
		3,
		3,
		1,
		3,
		2,
	}
	results := [][]int{
		{1, 2, 2, 3, 5, 6},
		{2, 3, 3, 4, 5, 6},
		{2, 2, 3, 4, 5, 6},
		{1},
		{0, 1, 1, 2, 3, 5},
		{0, 0, 1, 1, 2},
		{-111, -98, -3, -2, -1, 5, 6},
		{1, 2},
		{1, 2, 3, 4, 5, 6},
		{1, 1, 2},
	}
	for i := 0; i < len(testsNums1); i++ {
		merge2(testsNums1[i], testsM[i], testsNums2[i], testsN[i])
		if !reflect.DeepEqual(testsNums1[i], results[i]) {
			t.Fatalf("Wrong Answer, ret: %v right ret: %v", testsNums1[i], results[i])
		}
	}
}
