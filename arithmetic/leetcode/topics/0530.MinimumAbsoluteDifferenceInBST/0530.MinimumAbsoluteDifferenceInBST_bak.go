package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func inOrder2(node *TreeNode, preVal, minimum int) int {
	if node == nil {
		return minimum
	}

	minimum = inOrder2(node.Left, node.Val, minimum)

	if temp := abs(preVal, node.Val); temp < minimum {
		minimum = temp
	}

	minimum = inOrder2(node.Right, node.Val, minimum)

	return minimum
}

func getMinimumDifference2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return inOrder2(root, root.Val, 1<<16)
}
