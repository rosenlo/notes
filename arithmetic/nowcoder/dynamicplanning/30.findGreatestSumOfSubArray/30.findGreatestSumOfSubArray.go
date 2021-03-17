package main

/**
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	if len(array) == 0 {
		return 0
	}

	maximum := array[0]

	for i := 0; i < len(array); i++ {
		temp := array[i]
		for j := i + 1; j < len(array); j++ {
			temp += array[j]
			if temp > maximum {
				maximum = temp
			}
		}
		if temp > maximum {
			maximum = temp
		}
	}
	return maximum
}

func FindGreatestSumOfSubArray2(array []int) int {
	// write code here
	if len(array) == 0 {
		return 0
	}

	maximum, curSum := array[0], 0

	for i := 0; i < len(array); i++ {
		curSum += array[i]
		if array[i] > curSum {
			curSum = array[i]
		}
		if curSum > maximum {
			maximum = curSum
		}
	}
	return maximum
}
