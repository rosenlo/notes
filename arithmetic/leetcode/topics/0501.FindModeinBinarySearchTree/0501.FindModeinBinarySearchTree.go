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

func find(node *TreeNode, m map[int]int, max int) int {

	m[node.Val] += 1

	if m[node.Val] > max {
		max = m[node.Val]
	}

	if node.Left != nil {
		max = find(node.Left, m, max)
	}
	if node.Right != nil {
		max = find(node.Right, m, max)
	}
	return max
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	m := make(map[int]int)
	max := find(root, m, 0)
	size := len(m)

	ret := make([]int, 0, size)
	for nodeVal, freq := range m {
		if freq == max {
			ret = append(ret, nodeVal)
		}
	}
	return ret
}
