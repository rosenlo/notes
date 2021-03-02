package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := [][]int{
		[]int{0},
		[]int{9},
		[]int{1, 2, 3},
		[]int{4, 3, 2, 1},
		[]int{1, 9, 9},
		[]int{1, 0, 0, 0, 0},
	}
	results := [][]int{
		[]int{1},
		[]int{1, 0},
		[]int{1, 2, 4},
		[]int{4, 3, 2, 2},
		[]int{2, 0, 0},
		[]int{1, 0, 0, 0, 1},
	}
	for i := 0; i < len(tests); i++ {
		ret := plusOne(tests[i])
		if len(ret) != len(results[i]) {
			t.Errorf("Wrong length, ret: %v, result: %v", ret, results[i])
		}
		if !reflect.DeepEqual(ret, results[i]) {
			t.Errorf("Wrong Answer, ret: %v, result: %v", ret, results[i])
		}
	}
}
