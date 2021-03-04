# [559. Maximum Depth of N-ary Tree](https://leetcode.com/problems/maximum-depth-of-n-ary-tree/)

Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).



Example 1:

```
       1
     / | \
    3  2  4
   / \
  5   6

Input: root = [1,null,3,2,4,null,5,6]
Output: 3
```


Example 2:

```
               1
         /   |    |   \
        2    3    4     5
            / \   |    / \
           6   7  8   9   10
               |  |   |
               11 12  13
               |
               14


Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: 5
```

Constraints:

- The depth of the n-ary tree is less than or equal to 1000.
- The total number of nodes is between [0, 104].


### 题目大意

找到 N 叉数的最大深度

### 解题思路

和 [104. 二叉树的最大深度](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0104.MaximumDepthOfBinaryTree)方法一致，区别是二叉树是左右子树，N 叉数是子数组

### Accepted

Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Depth of N-ary Tree.

Memory Usage: 3.3 MB, less than 29.67% of Go online submissions for Maximum Depth of N-ary Tree.


### 总结

从二叉树到 N 叉数，本质没变化，无非子树路径多了一些

二叉树递归

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func travel(root *TreeNode) {
    travel(root.Left)
    travel(root.Right)
}

```

N 叉树递归

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func travel(root *TreeNode) {
    for _, node := range root.Children {
        travel(node)
    }
}

```
