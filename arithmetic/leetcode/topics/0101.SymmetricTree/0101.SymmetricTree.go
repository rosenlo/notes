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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func reversal(node *TreeNode) *TreeNode {
	if node == nil || node.Left == nil && node.Right == nil {
		return node
	}

	node.Left, node.Right = node.Right, node.Left

	reversal(node.Left)

	reversal(node.Right)

	return node
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	reversal(root.Left)

	return isSameTree(root.Left, root.Right)
}
