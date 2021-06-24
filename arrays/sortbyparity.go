package arrays

func sortArrayByParity(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	lo, hi := 0, len(nums)-1
	for ; lo < hi; {
		if nums[lo] % 2 != 0 && nums[hi] % 2 == 0 {
			nums[lo], nums[hi], lo, hi = nums[hi], nums[lo], lo+1, hi-1
		} else {
			if nums[lo] % 2 == 0 {
				lo++
			}
			if nums[hi] % 2 != 0 {
				hi--
			}
		}
	}
	return nums
}
