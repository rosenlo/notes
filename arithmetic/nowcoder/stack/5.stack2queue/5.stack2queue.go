package main

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack1) == 0 && len(stack2) == 0 {
		return -1
	}

	size := len(stack1)
	if len(stack2) == 0 {
		stack2 = make([]int, size)
		for i := 0; i < size; i++ {
			stack2[i] = stack1[size-i-1]
		}
		stack1 = []int{}
	}

	node := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return node
}
