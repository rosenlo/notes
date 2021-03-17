package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */
func rectCover(number int) int {
	// write code here
	if number < 4 {
		return number
	}

	return rectCover(number-1) + rectCover(number-2)
}
