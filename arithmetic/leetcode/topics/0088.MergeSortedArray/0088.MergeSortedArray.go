package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m+n != len(nums1) {
		return
	}

	if n != len(nums2) {
		return
	}

	if m < 0 || n > 200 {
		return
	}

	if m+n < 1 || m+n > 200 {
		return
	}

	if 0 == n {
		return
	}

	j, k := 0, 0
	for j < m {
		if nums2[k] < nums1[j] {
			temp := nums1[j]
			nums1[j] = nums2[k]
			nums2[k] = temp

			for i := 0; i+1 < n; i++ {
				if nums2[i] > nums2[i+1] {
					temp := nums2[i]
					nums2[i] = nums2[i+1]
					nums2[i+1] = temp
				}
			}
		}
		j++
	}
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	return
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1
	for m > 0 && n > 0 {
		if nums2[n-1] > nums1[m-1] {
			nums1[last] = nums2[n-1]
			n--
		} else {
			nums1[last] = nums1[m-1]
			m--
		}
		last--
	}
	for ; n > 0; n-- {
		nums1[last] = nums2[n-1]
		last--
	}
}
