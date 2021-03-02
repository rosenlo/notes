package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	test := []int{1, 3, 5, 6}
	targets := []int{
		5,
		2,
		7,
		0,
	}
	results := []int{
		2,
		1,
		4,
		0,
	}
	for i := 0; i < len(targets); i++ {
		if ret := searchInsert(test, targets[i]); ret != results[i] {
			t.Fatalf("Wrong Answer, target: %v ret: %v right ret: %v", targets[i], ret, results[i])
		}
	}
}
