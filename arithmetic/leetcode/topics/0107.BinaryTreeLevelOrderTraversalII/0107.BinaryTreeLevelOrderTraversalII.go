package leetcode

import "github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.TreeNode

func traversal(root *TreeNode) [][]int {
	ret := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodes := queue[0:]
		queue = []*TreeNode{}
		s := []int{}
		for _, node := range nodes {
			if node != nil {
				s = append(s, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		if len(s) != 0 {
			ret = append(ret, s)
		}
	}
	return ret
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret := traversal(root)

	size := len(ret)

	s := make([][]int, size)

	for i := size; i > 0; i-- {
		s[size-i] = ret[i-1]
	}
	return s
}
