package main

import "github.com/rosenlo/toolkits/structure/tree"

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

type TreeNode = tree.TreeNode

/**
 *
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	return abs(depthTravel(pRoot.Left), depthTravel(pRoot.Right)) <= 1 && IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)

}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depthTravel(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(depthTravel(node.Left), depthTravel(node.Right)) + 1
}
