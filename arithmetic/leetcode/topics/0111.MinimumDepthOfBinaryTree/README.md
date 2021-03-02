# [111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/)

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root
node down to the nearest leaf node.

**Note**: A leaf is a node with no children.

Example 1:

```
Input: root = [3,9,20,null,null,15,7]
Output: 2

       3
      / \
     9  20
       /  \
      15   7
```


Example 2:

```
Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5

       2
        \
          3
           \
            4
             \
              5
               \
                6

```

Constraints:

The number of nodes in the tree is in the range [0, 10^5].
-1000 <= Node.val <= 1000


### 题目大意

- 要求找到最小深度的叶子节点（无子节点）

### 解题思路

- 递归所有节点，只在叶子节点返回深度，然后取最小值
- 如果当前节点为空返回一个最大节点数，这里是10^5
