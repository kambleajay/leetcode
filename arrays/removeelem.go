package arrays

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, len(nums)-1
	for ; i < j; {
		if nums[i] == val && nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j-1
		} else if nums[i] != val {
			i++
		} else {
			j--
		}
	}
	if nums[i] == val {
		return i
	} else {
		return i + 1
	}
}
