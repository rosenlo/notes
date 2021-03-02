# [530. Minimum Absolute Difference in BST](https://leetcode.com/problems/minimum-absolute-difference-in-bst/)

Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

Example:

```
Input:

   1
    \
     3
    /
   2

Output:
1

Explanation:
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
```

Note:

- There are at least two nodes in this BST.
- This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/


### 题目大意

找到任意两个节点的最小差值

### 解题思路

暴力解法

- 利用二叉搜索树的特性转换成有序数组，然后两两相比较，返回最小值，但这样内存空间会使用很多

更优解法

- 利用二叉搜索树的特性用中序遍历比较相邻节点

### Accepted

Runtime: 12 ms, faster than 64.86% of Go online submissions for Minimum Absolute Difference in BST.

Memory Usage: 6.5 MB, less than 5.41% of Go online submissions for Minimum Absolute Difference in BST.
