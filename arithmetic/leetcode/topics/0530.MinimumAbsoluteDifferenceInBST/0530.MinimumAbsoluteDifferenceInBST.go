package leetcode

import (
	"github.com/rosenlo/notes/arithmetic/notes/arithmetic/leetcode/structure/tree"
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

func inOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append(inOrder(node.Left), node.Val), inOrder(node.Right)...)
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minimum := 1 << 16
	ints := inOrder(root)
	for i := 0; i < len(ints)-1; i++ {
		temp := ints[i+1] - ints[i]
		if temp < minimum {
			minimum = temp
		}
	}
	return minimum
}
