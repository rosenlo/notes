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


## 思路
- 找到数组中最大的元素 O(n)
- 找出大于零的元素开始折半查找，找到当前元素的最大区间
- 多个最大区间的最大值比较得出最终的最大区间、最大值

这一题可以用 DP 求解也可以不用 DP。
题目要求输出数组中某个区间内数字之和最大的那个值。dp[i] 表示 [0,i] 区间内各个子区间和的最大值，状态转移方程是 dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)，dp[i] = nums[i] (dp[i-1] ≤ 0)。

## testcase
- 数组中可能有多个最大值
- 总数不一定包含最大值
