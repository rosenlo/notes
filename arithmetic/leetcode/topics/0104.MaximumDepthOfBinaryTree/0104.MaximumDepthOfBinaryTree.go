package leetcode

import "github.com/rosenlo/notes/arithmetic/leetcode/structure/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.TreeNode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depthTravel(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	depth++

	return max(depthTravel(node.Left, depth), depthTravel(node.Right, depth))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	return depthTravel(root, depth)
}
