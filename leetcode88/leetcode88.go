package leetcode88

func merge(nums1 []int, m int, nums2 []int, n int) {
	for {
		if m == 0 || n == 0 {
			break
		}

		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m = m - 1
		} else {
			nums1[m+n-1] = nums2[n-1]
			n = n - 1
		}
	}

	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n = n - 1
	}
}
