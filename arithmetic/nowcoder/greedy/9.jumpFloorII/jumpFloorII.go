package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */

func pow(x, n int) int {
	ret := 1
	for n != 0 {
		ret *= x
		n--
	}
	return ret
}

func jumpFloorII(number int) int {
	// write code here

	if number < 3 {
		return number
	}
	return pow(2, number-1)
}
