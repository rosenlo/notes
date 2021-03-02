# [1.Two Sum](https://leetcode.com/problems/two-sum/)

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 思路

新建一个map，目标值循环相减数组元素，如果差值在 map 中则返回数组下标，否则放入 map

- 时间复杂度 O(n) 需要遍历一次数组
- 空间复杂度 O(n) map 依赖数组元素数量
