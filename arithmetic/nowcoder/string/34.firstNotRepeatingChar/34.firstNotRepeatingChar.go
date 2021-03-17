package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return int整型
 */
func FirstNotRepeatingChar(str string) int {
	// write code here

	if len(str) == 0 {
		return -1
	}

	m := make(map[uint8]int)
	for i := 0; i < len(str); i++ {
		m[str[i]] += 1
	}
	for i := 0; i < len(str); i++ {
		if m[str[i]] == 1 {
			return i
		}
	}
	return -1
}
