# [第一个只出现一次的字符](https://www.nowcoder.com/practice/1c82e8cf713b4bbeb2a5b31cf5b0417c?tpId=13&tqId=11187&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）

## 示例1

```
输入
"google"

返回值
4
```

## 解题思路

第一时间想到的还是用 hash 表，记录出现的次数。这种方式可行但不是最优解

先快速刷第一遍，第二遍再做优化

## Accepted

运行时间：3ms 超过22.54%用Go提交的代码

占用内存：824KB 超过67.61%用Go提交的代码
