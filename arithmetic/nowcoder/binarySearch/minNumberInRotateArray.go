package main

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 0 {
		return 0
	}

	minimum := rotateArray[len(rotateArray)-1]
	for i := len(rotateArray); i > 0; i-- {
		if rotateArray[i-1] < minimum {
			minimum = rotateArray[i-1]
		}
	}

	return minimum
}
