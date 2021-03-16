# [不用加减乘除做加法](https://www.nowcoder.com/practice/59ac416b4b944300b617d4f7f111b215?tpId=13&tqId=11201&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

写一个函数，求两个整数之和，要求在函数体内不得使用+、-、\*、/四则运算符号。

## 示例1

```
输入
1,2

返回值
3
```

## 分类

数学

## 解题思路

a ^ b 表示没有考虑进位的情况下两数的和，(a & b) << 1 就是进位。

递归会终止的原因是 (a & b) << 1 最右边会多一个 0，那么继续递归，进位最右边的 0 会慢慢增多，最后进位会变为 0，递归终止。


## 总结

这题考虑到了用或运算，但没考虑到进位的情况。
