package main

import (
	"testing"
)

type param struct {
	one    int
	result int
}

func TestKthNode(t *testing.T) {
	tests := []param{
		{
			4,
			5,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := rectCover(tests[i].one)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret, tests[i].result)
		}
	}
}
