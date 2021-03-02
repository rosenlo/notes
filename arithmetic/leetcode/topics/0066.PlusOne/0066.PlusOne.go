package leetcode

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 0 {
			continue
		}
		if digits[i]%10 == 0 {
			digits[i] = 0
			if i != 0 {
				digits[i-1]++
			} else {
				newDigits := []int{1}
				newDigits = append(newDigits, digits...)
				return newDigits
			}
		}
	}
	return digits
}
