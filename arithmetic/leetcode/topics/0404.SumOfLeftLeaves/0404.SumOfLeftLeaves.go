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

func travelLeftLeaves(node *TreeNode, left bool, sum int) int {
	if left && node.Left == nil && node.Right == nil {
		return sum + node.Val
	}

	if node.Left != nil {
		sum = travelLeftLeaves(node.Left, true, sum)
	}
	if node.Right != nil {
		sum = travelLeftLeaves(node.Right, false, sum)
	}
	return sum
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root != nil {
		return travelLeftLeaves(root, false, 0)
	}
	return 0
}
