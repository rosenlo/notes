package leetcode

import (
	"github.com/rosenlo/notes/arithmetic/leetcode/structure/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.TreeNode

func depthTravel(node *TreeNode, valueSum, targetSum int) bool {
	if node == nil {
		return false
	}
	valueSum += node.Val

	if node.Left == nil && node.Right == nil {
		return valueSum == targetSum
	}

	return depthTravel(node.Left, valueSum, targetSum) || depthTravel(node.Right, valueSum, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return depthTravel(root, 0, targetSum)
}
