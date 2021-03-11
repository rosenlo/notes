package main

import (
	"fmt"
	"testing"
)

type param struct {
	one    string
	result int
}

func TestFristNotRepeatingChar(t *testing.T) {
	tests := []param{
		{
			"google",
			4,
		},
		{
			"googgle",
			5,
		},
		{
			"abcd",
			0,
		},
		{
			"aabb",
			-1,
		},
		{
			"NXWtnzyoHoBhUJaPauJaAitLWNMlkKwDYbbigdMMaYfkVPhGZcrEwp",
			1,
		},
	}
	for i := 0; i < len(tests); i++ {
		ret := FirstNotRepeatingChar(tests[i].one)
		if ret != tests[i].result {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i].one, ret, tests[i].result)
		}
		fmt.Printf("[input]: %v\t", tests[i].one)
		fmt.Printf("[expect]: %v\t", tests[i].result)
		fmt.Printf("[output]: %v\t\n", ret)
	}

}
