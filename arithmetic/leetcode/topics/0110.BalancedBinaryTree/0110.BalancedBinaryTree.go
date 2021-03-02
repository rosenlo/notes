package leetcode

import "github.com/rosenlo/leetcode/structure/tree"

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

func abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func depthTravel(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	depth++

	return max(depthTravel(node.Left, depth), depthTravel(node.Right, depth))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(depthTravel(root.Left, 1), depthTravel(root.Right, 1)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
