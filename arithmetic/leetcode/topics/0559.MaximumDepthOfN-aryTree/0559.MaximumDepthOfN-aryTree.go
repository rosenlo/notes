package bst

import (
	"github.com/rosenlo/notes/arithmetic/leetcode/structure/tree"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node = tree.Node

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depthTravel(root *Node, depth int, maximum *int) int {
	if root == nil {
		return depth
	}
	for _, node := range root.Children {
		*maximum = max(*maximum, depthTravel(node, depth+1, maximum))
	}

	return depth

}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maximum := 1
	depthTravel(root, maximum, &maximum)
	return maximum
}
