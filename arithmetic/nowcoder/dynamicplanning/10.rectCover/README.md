# [矩形覆盖](https://www.nowcoder.com/practice/72a5a919508a4251859fb2cfb987a0e6?tpId=13&tqId=11163&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

我们可以用2 * 1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2 * 1的小矩形无重叠地覆盖一个2 * n的大矩形，总共有多少种方法？

比如n = 3 时，2 * 3的矩形块有3种覆盖方法：


```
-------
| | | |
| | | |
-------

--------
| |    |
| |----|
| |    |
--------

--------
|    | |
|----| |
|    | |
--------
```

## 示例1

```
输入
4

返回值
5
```

## 知识点

递归


## 解题思路

这题和 [9.jumpFloor](https://github.com/rosenlo/notes/tree/master/arithmetic/nowcoder/greedy/8.jumpFloor) 基本一致，也是 fibonacci 数列， 利用递归来解决
