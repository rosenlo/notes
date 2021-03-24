package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pushV int整型一维数组
 * @param popV int整型一维数组
 * @return bool布尔型
 */

func IsPopOrder(pushV []int, popV []int) bool {
	// write code here
	if len(pushV) != len(popV) {
		return false
	}

	outSeq := 0
	stack := make([]int, 0, len(pushV))
	for i := 0; i < len(pushV); i++ {
		if pushV[i] == popV[outSeq] {
			outSeq++
		} else {
			stack = append(stack, pushV[i])
		}
		for len(stack) > 0 {
			if stack[len(stack)-1] != popV[outSeq] {
				break
			}
			outSeq++
			stack = stack[:len(stack)-1]
		}
	}
	if outSeq == len(pushV) {
		return true
	}
	return false

}
