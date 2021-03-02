package leetcode

import (
	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.TreeNode

func commonAncestor(node, p, q *TreeNode) *TreeNode {
	if p != nil && q != nil {
		return node
	}
	if p != nil {
		return p
	}
	return q
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	return commonAncestor(root, lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q))

}
