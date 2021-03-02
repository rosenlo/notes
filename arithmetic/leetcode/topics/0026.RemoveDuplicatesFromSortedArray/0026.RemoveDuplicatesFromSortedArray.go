package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	preIndex := 0
	preValue := nums[preIndex]
	for i := 0; i < len(nums); i++ {
		if nums[i] == preValue {
			continue
		}
		preIndex++
		preValue = nums[i]
		nums[preIndex] = preValue
	}
	return preIndex + 1
}
