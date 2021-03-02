package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for idx := 0; idx < len(nums); idx++ {
		tempSum, tempMax := nums[idx], nums[idx]
		for i := idx - 1; i >= 0; i-- {
			tempSum += nums[i]
			if tempSum > tempMax {
				tempMax = tempSum
			}
		}
		tempSum = nums[idx]
		for i := idx + 1; i < len(nums); i++ {
			tempSum += nums[i]
			if tempSum > tempMax {
				tempMax = tempSum
			}

		}
		if tempMax > maxVal {
			maxVal = tempMax
		}
	}
	return maxVal
}

func maxSubArray2(nums []int) int {
	maxSum, curSum := -2147483647, 0
	for i := 0; i < len(nums); i++ {
		curSum = curSum + nums[i]
		if nums[i] > curSum {
			curSum = nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
