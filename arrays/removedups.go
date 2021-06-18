package arrays

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	n, src, dst := nums[0], 1, 1
	for ;src < len(nums); {
		if nums[src] == n {
			src++
		} else {
			nums[dst], dst, src, n = nums[src], dst+1, src+1, nums[src]
		}
	}
	return dst
}
