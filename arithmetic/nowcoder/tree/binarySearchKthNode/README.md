# [二叉搜索树的第k个节点](https://www.nowcoder.com/practice/ef068f602dde4d28aab2b210e859150a?tpId=13&tqId=11215&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

给定一棵二叉搜索树，请找出其中的第k小的TreeNode结点。

## 示例1

```
输入
{5,3,7,2,4,6,8},3

返回值
{4}
```

## 说明

按结点数值大小顺序第三小结点的值为4


## 解题思路

利用二叉搜索树的特性，中序遍历出排第k位的节点
