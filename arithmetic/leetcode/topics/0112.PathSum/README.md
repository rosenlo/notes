# [112. Path Sum](https://leetcode.com/problems/path-sum/)

Given the `root` of a binary tree and an integer `targetSum`, return `true` if
the tree has a **root-to-leaf** path such that adding up all the values along
the path equals `targetSum`.

A **leaf** is a node with no children.

**Example 1:**

```
         5
        / \
       4   8
      /   / \
     11  13  4
    / \       \
   7   2       1

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
```

**Example 2:**

```
         1
        / \
       2   3

Input: root = [1,2,3], targetSum = 5
Output: false
```

**Example 3:**

```
         1
        /
       2

Input: root = [1,2], targetSum = 0
Output: false
```

Constraints:

- The number of nodes in the tree is in the range `[0, 5000]`.
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

### 题目大意

- 要求找到根节点到叶子节点（无子节点）的路径和与给定的 `targetSum` 是否匹配

### 解题思路

- 与 [0111.Minimum Depth of Binary Tree](https://github.com/rosenlo/leetcode/tree/master/topics/0111.MinimumDepthOfBinaryTree) 基本一致，不同的是把深度换成路径和长度
