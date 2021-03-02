# [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)

Invert a binary tree.

Example:

```
Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1

```

Trivia:
This problem was inspired by this original tweet by Max Howell:

> Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off...

### 题目大意

- 反转一颗二叉树

### 解题思路

- 与[101.Symmetric Tree](https://github.com/rosenlo/leetcode/tree/master/topics/0101.SymmetricTree) 类似，利用递归同时反转左右子节点
