package main

/**
 *
 * @param n int整型
 * @return int整型
 */

func Fibonacci(n int) int {
	// write code here
	if n > 39 {
		return -1
	}
	if n == 0 || n == 1 {
		return n
	}
	s := make([]int, 40)
	s[0] = 0
	s[1] = 1
	for i := 2; i <= 39; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]
}
