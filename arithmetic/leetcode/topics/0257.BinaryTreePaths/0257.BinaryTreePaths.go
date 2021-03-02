package leetcode

import (
	"fmt"
	"strconv"

	"github.com/rosenlo/notes/arithmetic/leetcode/structure/tree"
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

func travel(node *TreeNode, paths []string, path string) []string {
	if node == nil {
		return paths
	}

	if path != "" {
		path = fmt.Sprintf("%s->%d", path, node.Val)
	} else {
		path = fmt.Sprintf("%d", node.Val)
	}

	if node.Left == nil && node.Right == nil {
		paths = append(paths, path)
	}

	paths = travel(node.Left, paths, path)

	paths = travel(node.Right, paths, path)

	return paths

}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	return travel(root, []string{}, "")

}

/*

Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Paths.
Memory Usage: 2.4 MB, less than 44.68% of Go online submissions for Binary Tree Paths.

*/
func travel2(node *TreeNode, paths []string, path string) []string {
	if node.Left == nil && node.Right == nil {
		paths = append(paths, path+strconv.Itoa(node.Val))
	}

	if node.Left != nil {
		paths = travel2(node.Left, paths, path+strconv.Itoa(node.Val)+"->")
	}

	if node.Right != nil {
		paths = travel2(node.Right, paths, path+strconv.Itoa(node.Val)+"->")
	}

	return paths
}

func binaryTreePaths2(root *TreeNode) []string {
	if root != nil {
		return travel2(root, []string{}, "")
	}
	return []string{}

}
