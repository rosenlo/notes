# [构建乘积数组](https://www.nowcoder.com/practice/94a4d381a68b47b7a8bed86f2975db46?tpId=13&tqId=11204&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

## 题目描述

给定一个数组 A[0,1,...,n-1]， 请构建一个数组 B[0,1,...,n-1]， 其中B 中的元素 B[i]=A[0] * A[1] * ... * A[i-1] * A [i+1] * ... *  A [n-1]。不能使用除法。

注意：规定B[0] = A[1] * A[2] * ... * A[n-1]，B[n-1] = A[0] * A[1] * ... * A[n-2];） 对于A长度为1的情况，B无意义，故而无法构建，因此该情况不会存在。

## 示例1

```
输入
[1,2,3,4,5]

返回值
[120,60,40,30,24]
```

## 题目大意

B 乘积数组的意思是指在数组 A 中，除 i 以外的元素相乘，比如： B[1] = 1 * 3 * 4 * 5 = 60
