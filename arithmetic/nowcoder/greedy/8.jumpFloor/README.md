# [跳台阶](https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4?tpId=13&tqId=11161&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。


##  解题思路

f(1) = 1

f(2) = 2

f(3) = 3

f(4) = 5

f(5) = 8

f(6) = 13

f(n) = f(n-1) + f(n-2); n > 3

从规律上来看是 fibonacci 数列

