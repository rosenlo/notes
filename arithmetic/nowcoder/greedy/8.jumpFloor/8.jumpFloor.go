package main

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here

	if number < 4 {
		return number
	}

	return jumpFloor(number-1) + jumpFloor(number-2)
}
