# [0100. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

> a binary tree in which the left and right **subtrees** of every node differ in height by no more than 1.


Example 1:

```
Input: root = [3,9,20,null,null,15,7]
Output: true

     3
    / \
   9  20
     /  \
    15   7
```

Example 2:

```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

         1
        / \
       2   2
      / \
     3   3
    / \
   4   4
```

Example 3:

```
Input: root = []
Output: true
```

Constraints:

The number of nodes in the tree is in the range [0, 5000].
-10^4 <= Node.val <= 10^4

## 思路

高度平衡的二叉树定义为：其每个节点的左右子树高度相差不超过 `1`，可得知以下几点：

- 每个节点
- 左右子树高度 `<= 1`

即可递归每个节点，判断当前节点的左右子树高度是否小于等于 `1`
