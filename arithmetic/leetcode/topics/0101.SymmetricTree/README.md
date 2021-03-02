# [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)

Given a binary tree, check whether it is a mirror of itself (ie, symmetric
around its center).

For example, this binary tree `[1, 2, 2, 3, 4, 4, 3]` is symmetric.

```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

But the following `[1, 2, 2, null ,3, null, 3]` is not:

```
    1
   / \
  2   2
   \   \
   3    3
```

**Follow up**: Solve it both recursively and iteratively.

## 思路

- 对称的含义为：图形或物体相对的两边各部分，在大小、形状、距离和排列等方面一一相当
- 对称树则可以围绕着中心节点（root）将一半子树对称折叠（反转）再进行 2 棵子树对比（same tree）
