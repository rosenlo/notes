# [连续子数组的最大和](https://www.nowcoder.com/practice/459bd355da1549fa8a49e350bf3df484?tpId=13&tqId=11183&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为 O(n).

### 示例1

```
输入
[1,-2,3,10,-4,7,2,-5]

返回值
18
```

### 说明

输入的数组为{1,-2,3,10,—4,7,2,-5}，和最大的子数组为{3,10,-4,7,2}，因此输出为该子数组的和 18。

### 类型

动态规划、分治


## 解题思路

与 LeetCode 上的 [53. Maximum Subarray](https://github.com/rosenlo/notes/tree/master/arithmetic/leetcode/topics/0053.MaximumSubarray)  一致


## 总结

在迭代循环中拿当前值和之前的累加值比较，如果大于之前的累加值，说明当前值为临时最大值，如果当前值大于最大和，最大和等于当前值

核心的区别是不需要把后面的元素全部加完，迭代到当前整数比之前的累计值大的话就可以分配临时最大值了，也就不需要两层循环做这个事了
