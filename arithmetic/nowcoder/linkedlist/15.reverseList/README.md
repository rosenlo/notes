# [反转链表](https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca?tpId=13&tqId=11168&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

输入一个链表，反转链表后，输出新链表的表头。

```
示例1
输入

{1,2,3}
返回值
{3,2,1}
```

## 解题思路

需要使用一个辅助空间 `preNode` 保存前置节点，迭代链表进行交换即可

空间复杂度：O(1)
时间复杂度：O(n)
