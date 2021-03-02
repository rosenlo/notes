package tree

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
