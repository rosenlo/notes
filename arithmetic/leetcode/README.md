# LeetCode

**只记录已完成的题**

No.|Title|Solution|Acceptance|Difficulty|
:--:|:-----:|:--------:|:----------:|:----------:|
0001|Two Sum|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0001.TwoSum)|46.4%|Easy
0026|Remove Duplicates From Sorted Array|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0026.RemoveDuplicatesFromSortedArray)|46.5%|Easy
0027|Remove Element|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0027.RemoveElement)|49.2%|Easy
0035|Search Insert Position|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0035.SearchInsertPosition)|42.8%|Easy
0053|Maximum Subarray|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0053.MaximumSubarray)|47.7%|Easy
0066|Plus One|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0066.PlusOne)|42.4%|Easy
0088|Merge Sorted Array|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0088.MergeSortedArray)|40.7%|Easy
0100|Same Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0100.SameTree)|54.1%|Easy
0101|Symmetric Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0101.SymmetricTree)|48.0%|Easy
0104|Maximum Depth Of Binary Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0104.MaximumDepthOfBinaryTree)|67.9%|Easy
0107|Binary Tree Level Order Traversal II|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0107.BinaryTreeLevelOrderTraversalII)|55.1%|Medium
0108|Convert Sorted Array To Binary Search Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0108.ConvertSortedArrayToBinarySearchTree)|60.2%|Easy
0110|Balanced Binary Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0110.BalancedBinaryTree)|44.7%|Easy
0111|Minimum Depth of Binary Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0111.MinimumDepthOfBinaryTree)|39.5%|Easy
0112|Path Sum|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0112.PathSum)|42.4%|Easy
0226|Invert Binary Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0226.InvertBinaryTree)|66.9%|Easy
0235|Lowest Common Ancestor of a Binary Search Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0235.LowestCommonAncestorOfABinarySearchTree)|51.7%|Easy
0257|Binary Tree Paths|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0257.BinaryTreePaths)|53.6%|Easy
0404|Sum of Left Leaves|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0404.SumOfLeftLeaves)|52.3%|Easy
0501|Find Mode in Binary Search Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0501.FindModeinBinarySearchTree)|43.4%|Easy
0530|Minimum Absolute Difference in BST|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0530.MinimumAbsoluteDifferenceInBST)|54.8%|Easy
0543|Diameter of Binary Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0543.DiameterOfBinaryTree)|49.1%|Easy
0559|Maximum Depth of N-ary Tree|[Go](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0559.MaximumDepthOfN-aryTree)|69.5%|Easy

## 总结

### Tree

#### 二叉树遍历

```go
/*
      1
     / \
    2   3
   / \   \
  4   5   6

  层次遍历：[1, 2, 3, 4, 5, 6]
  前序遍历：[1, 2, 4, 5, 3, 6]
  中序遍历: [4, 2, 5, 1, 3, 6]
  后序遍历：[4, 5, 2, 6, 3, 1]

*/

// 前序遍历
func PreOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append([]int{node.Val}, PreOrder(node.Left)...), PreOrder(node.Right)...)
}

// 中序遍历
func InOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append(InOrder(node.Left), node.Val), InOrder(node.Right)...)
}

// 后序遍历
func PostOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append(PostOrder(node.Left), PostOrder(node.Right)...), node.Val)
}
```

#### 二叉树递归

```go
/*
  Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
*/

func travel(root *TreeNode) {
    travel(root.Left)
    travel(root.Right)
}

```

#### N 叉树递归

```go
/*
  Definition for a Node.
  type Node struct {
      Val int
      Children []*Node
  }
*/

func travel(root *TreeNode) {
    for _, node := range root.Children {
        travel(node)
    }
}

```
