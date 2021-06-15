package arrays

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; k > -1; k-- {
		nums1HasElems := i > -1
		nums2HasElems := j > -1
		if !nums2HasElems {
			break
		} else if !nums1HasElems {
			copy(nums1[:k+1], nums2[:j+1])
			break
		} else if nums1[i] > nums2[j] {
			nums1[k], i = nums1[i], i-1
		} else {
			nums1[k], j = nums2[j], j-1
		}
	}
}
