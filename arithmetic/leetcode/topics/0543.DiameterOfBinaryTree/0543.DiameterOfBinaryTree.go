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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
Runtime: 16 ms, faster than 8.07% of Go online submissions for Diameter of Binary Tree.

Memory Usage: 4.9 MB, less than 8.07% of Go online submissions for Diameter of Binary Tree.
*/

func depthTravel(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	return max(depthTravel(node.Left, depth), depthTravel(node.Right, depth)) + 1

}

func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maximum := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			if temp := depthTravel(node.Left, 0) + depthTravel(node.Right, 0); temp > maximum {
				maximum = temp
			}
			queue = append(queue, node.Left, node.Right)
		}

	}

	return maximum
}

/**
Runtime: 4 ms, faster than 96.86% of Go online submissions for Diameter of Binary Tree.

Memory Usage: 4.5 MB, less than 56.50% of Go online submissions for Diameter of Binary Tree.
*/

func depthTravel2(node *TreeNode, maximum *int) int {
	if node == nil {
		return 0
	}

	left, right := depthTravel2(node.Left, maximum), depthTravel2(node.Right, maximum)

	*maximum = max(*maximum, left+right)

	return max(left, right) + 1

}

func diameterOfBinaryTree2(root *TreeNode) int {

	if root == nil {
		return 0
	}
	var maximum int
	depthTravel2(root, &maximum)
	return maximum
}
