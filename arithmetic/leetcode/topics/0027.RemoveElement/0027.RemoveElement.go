package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	preIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[preIndex] = nums[i]
		preIndex++
	}
	return preIndex
}
