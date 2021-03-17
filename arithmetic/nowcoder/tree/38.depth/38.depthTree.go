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
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot TreeNode类
 * @return int整型
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TreeDepth(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}

	return max(TreeDepth(pRoot.Left), TreeDepth(pRoot.Right)) + 1
}
