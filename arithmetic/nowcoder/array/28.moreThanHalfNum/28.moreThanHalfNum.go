package main

/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here

	if len(numbers) == 0 {
		return 0
	}

	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		m[numbers[i]] += 1
	}

	maximum := 0
	ret := 0
	for k, v := range m {
		if v > maximum {
			maximum = v
			ret = k
		}
	}

	if maximum > len(numbers)/2 {
		return ret
	}

	return 0
}
