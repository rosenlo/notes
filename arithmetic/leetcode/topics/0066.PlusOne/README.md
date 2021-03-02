# [66. Plus One](https://leetcode.com/problems/plus-one/)
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

```
Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

Example 2:

```
Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```

## 思路
- 个位加一（数组最后一位）
- 迭代数组
- 逢十进一（非`0`值与`10`取余）
- 如果最高位需要进一需要 insert 到数组首位或生成一个新数组
