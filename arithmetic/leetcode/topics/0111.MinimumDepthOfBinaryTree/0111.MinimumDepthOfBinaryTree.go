package leetcode

import (
	"math"

	"github.com/rosenlo/leetcode/structure/tree"
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

var MaxNodes int = int(math.Pow(10, 5))

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func depthTravel(node *TreeNode, depth int) int {
	if node == nil {
		return MaxNodes
	}

	depth++

	if node.Left == nil && node.Right == nil {
		return depth
	}

	return min(depthTravel(node.Left, depth), depthTravel(node.Right, depth))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return depthTravel(root, 0)
}
