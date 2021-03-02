[543. Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the **longest** path between any two nodes in a tree. This path may or may not pass through the root.

Example:

```
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
```

Return **3**, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.


### 题目大意

任意给两个节点，计算两个节点最长的直径距离

### 解题思路

任意两个节点的最长路径，就是两个节点到中心节点的深度相加

### Accepcted

- 第一次提交（遍历）

Runtime: 16 ms, faster than 8.07% of Go online submissions for Diameter of Binary Tree.

Memory Usage: 4.9 MB, less than 8.07% of Go online submissions for Diameter of Binary Tree.


- 第二次提交（迭代）

Runtime: 4 ms, faster than 96.86% of Go online submissions for Diameter of Binary Tree.

Memory Usage: 4.5 MB, less than 56.50% of Go online submissions for Diameter of Binary Tree.


### 总结

第一次提交虽然掌握了解题的思路，但是方式不是最优解，没意识到 `maximum`
可以用指针传递到 `depthTravel` 里，只注重了函数的 `return` 只能有一个
`depth`，其实return 两个也行，但内存空间会使用多一些。陷入了这个死角。

另外内存空间从 Rank 来看还有很大的优化空间

时间复杂度：O(n)
空间复杂度：O(n)
