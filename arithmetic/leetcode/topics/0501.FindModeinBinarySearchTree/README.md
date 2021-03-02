# [501. Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)

Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.

Assume a BST is defined as follows:

- The left subtree of a node contains only nodes with keys less than or equal to the node's key.
- The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
- Both the left and right subtrees must also be binary search trees.


For example:

```
Given BST [1,null,2,2],

   1
    \
     2
    /
   2

return [2].

```

Note: If a tree has more than one mode, you can return them in any order.

Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).

### 题目大意

在二叉搜索树中找到最常出现的元素

### 解题思路

暴力解法

- 找到出现次数最多的元素O(n)，然后把相同次数的元素找出来O(n)

更优解法

- 减少循环次数，减少额外空间的使用


### Accepted

Runtime: 12 ms, faster than 67.14% of Go online submissions for Find Mode in Binary Search Tree.

Memory Usage: 6.3 MB, less than 27.14% of Go online submissions for Find Mode in Binary Search Tree.
