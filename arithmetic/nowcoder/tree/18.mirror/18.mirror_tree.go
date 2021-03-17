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
 *
 * @param pRoot TreeNode类
 * @return TreeNode类
 */

func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return pRoot
	}
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left

	Mirror(pRoot.Left)
	Mirror(pRoot.Right)

	return pRoot

}
