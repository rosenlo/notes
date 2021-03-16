package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param A int整型一维数组
 * @return int整型一维数组
 */

func multiply(A []int) []int {
	// write code here
	size := len(A)
	B := make([]int, size)

	for i := 0; i < size; i++ {
		product := 1
		for j := 0; j < size; j++ {
			if j != i {
				product *= A[j]
			}
		}
		B[i] = product
	}
	return B
}
