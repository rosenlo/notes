# [数组中出现次数超过一半的数字](https://www.nowcoder.com/practice/e8a1b01a2df14cb2b228b30ee6a92163?tpId=13&tqId=11181&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。如果不存在则输出0。


示例1

```
输入
[1,2,3,2,2,2,5,4,2]

返回值
2
```

## 解题思路

题目分类归为 `hash` ，明示了可以利用 `hash` 表来解决，不过这种方式需要额外的内存空间，时间复杂度也比较高，感觉不是最优解，先 AC 后有时间再优化。
