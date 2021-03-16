package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param num1 int整型
 * @param num2 int整型
 * @return int整型
 */

func Add(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	}
	return Add(num1^num2, (num1&num2)<<1)
}
