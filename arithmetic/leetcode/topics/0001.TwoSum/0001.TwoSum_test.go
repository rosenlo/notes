package leetcode

import "testing"

var (
	tests = [][]int{
		[]int{2, 5, 11, 15},
		[]int{3, 3},
		[]int{0, 1, 4, 6, 9, 12},
	}

	targets = []int{
		16,
		6,
		9,
	}

	result = [][]int{
		[]int{1, 2},
		[]int{0, 1},
		[]int{0, 4},
	}
)

func TestTwoSum(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		if ret := twoSum(tests[i], targets[i]); ret[0] != result[i][0] && ret[1] != result[i][1] {
			t.Fatalf("Wrong Answer, ret: %v result: %v", ret, result[i])
		}
	}
}

func TestTwoSum2(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		if ret := twoSum2(tests[i], targets[i]); ret[0] != result[i][0] && ret[1] != result[i][1] {
			t.Fatalf("Wrong Answer, ret: %v result: %v", ret, result[i])
		}
	}
}

func TestTwoSum3(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		if ret := twoSum3(tests[i], targets[i]); ret[0] != result[i][0] && ret[1] != result[i][1] {
			t.Fatalf("Wrong Answer, ret: %v result: %v", ret, result[i])
		}
	}
}
