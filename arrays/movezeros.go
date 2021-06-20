package arrays

func moveZeroes(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	dst, src := 0, 0
	for ;src < len(nums); {
		if nums[src] != 0 {
			nums[dst], dst, src = nums[src], dst+1, src+1
		} else {
			src++
		}
	}
	for i := dst; i < len(nums); i++ {
		nums[i] = 0
	}
}
