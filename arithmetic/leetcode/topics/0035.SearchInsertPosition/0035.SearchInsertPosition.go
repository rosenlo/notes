package leetcode

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}
