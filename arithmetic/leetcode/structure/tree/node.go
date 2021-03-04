package tree

// Node represented as N-ary Tree Node
type Node struct {
	Val      int
	Children []*Node
}

// TreeNode represented as Binary Tree Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = -1 << 16

/**
      1
     / \
    2   3
   / \   \
  4   5   6

  层次遍历：[1, 2, 3, 4, 5, 6]
  前序遍历：[1, 2, 4, 5, 3, 6]
  中序遍历: [4, 2, 5, 1, 3, 6]
  后续遍历：[4, 5, 2, 6, 3, 1]

*/
func PreOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append([]int{node.Val}, PreOrder(node.Left)...), PreOrder(node.Right)...)
}

func InOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append(InOrder(node.Left), node.Val), InOrder(node.Right)...)
}

func PostOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append(PostOrder(node.Left), PostOrder(node.Right)...), node.Val)
}

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	for i := 1; i < n; i++ {
		node := queue[0]
		queue = queue[1:]

		if ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}

		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
	}

	return root
}

func TreeNode2Ints(root *TreeNode) []int {
	ints := make([]int, 0)
	if root == nil {
		return ints
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			if len(queue) > 0 {
				ints = append(ints, NULL)
			}
		} else {
			ints = append(ints, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	n := len(ints)
	for n > 0 && ints[n-1] == NULL {
		n--
	}
	return ints[:n]
}

/**

N-ary Tree

1. 输出序列以层级表示，每个子树组通过 NULL 值分割
2. 存在子树为 NULL 值的情况，如果遇到 NULL 值，表明当前节点无子数组

*/
func Ints2NaryTreeNode(ints []int) *Node {
	size := len(ints)
	if size == 0 {
		return nil
	}

	root := &Node{Val: ints[0]}

	if size == 1 {
		return root
	}

	queue := []*Node{root}

	children := make([]*Node, 0, size)

	var node *Node
	for i := 2; i < size; i++ {
		node = queue[0]

		if ints[i] == NULL {
			if len(children) > 0 {
				node.Children = append(node.Children, children...)
				queue = append(queue[1:], children...)
				children = []*Node{}
			} else {
				queue = queue[1:]
			}
		} else {
			children = append(children, &Node{Val: ints[i]})
		}

	}
	if len(children) != 0 {
		node.Children = children
	}

	return root
}

func NaryTreeNode2Ints(node *Node) []int {
	if node == nil {
		return []int{}
	}
	ret := []int{node.Val}
	queue := []*Node{node}

	if len(node.Children) > 0 {
		ret = append(ret, NULL)
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, children := range node.Children {
			ret = append(ret, children.Val)
		}
		if len(node.Children) != 0 {
			queue = append(queue, node.Children...)
		}
		if len(queue) > 0 {
			ret = append(ret, NULL)
		}
	}
	n := len(ret)

	for n > 0 && ret[n-1] == NULL {
		n--
	}

	return ret[:n]
}
