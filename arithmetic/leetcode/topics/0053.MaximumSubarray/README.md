# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。


Example:

```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.


## 解题思路

在迭代循环中拿当前值和之前的累加值比较，如果大于之前的累加值，说明当前值为临时最大值，如果当前值大于最大和，最大和等于当前值

核心的区别是不需要把后面的元素全部加完，迭代到当前整数比之前的累计值大的话就可以分配临时最大值了，也就不需要两层循环做这个事了
