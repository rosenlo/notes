package main

import (
	"github.com/rosenlo/toolkits/structure/tree"
)

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
 * @param k int整型
 * @return TreeNode类
 */

func KthNode(pRoot *TreeNode, k int) *TreeNode {
	// write code here

	if pRoot == nil {
		return nil
	}

	cur := 0
	return inOrder(pRoot, &cur, k)
}

func inOrder(node *TreeNode, cur *int, k int) *TreeNode {
	if node == nil {
		return nil
	}

	if n := inOrder(node.Left, cur, k); n != nil {
		return n
	}

	*cur++
	if *cur == k {
		return node
	}

	return inOrder(node.Right, cur, k)
}
