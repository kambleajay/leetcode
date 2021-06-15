package arrays

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; k > -1; k-- {
		if i > -1 && j > -1 {
			if nums1[i] > nums2[j] {
				nums1[k], i = nums1[i], i-1
			} else {
				nums1[k], j = nums2[j], j-1
			}
		} else if i > -1 {
			nums1[k], i = nums1[i], i-1
		} else {
			nums1[k], j = nums2[j], j-1
		}
	}
}
