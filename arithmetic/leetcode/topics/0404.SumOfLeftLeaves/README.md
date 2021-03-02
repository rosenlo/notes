# [404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/)

Find the sum of all left leaves in a given binary tree.

Example:

```
    3
   / \
  9  20
    /  \
   15   7
```

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.


### 题目大意

找到所有左子树的叶子节点累加和

### 解题思路

还是使用递归，关键点是如何识别左子树，这里的实现是加了一个标识符，开销就是额外的内存空间。

### Accepted

Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Left Leaves.

Memory Usage: 2.7 MB, less than 22.58% of Go online submissions for Sum of Left Leaves.
