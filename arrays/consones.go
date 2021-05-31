package arrays

//findMaxConsecutiveOnes returns the maximum number of consecutive 1s in nums.
func findMaxConsecutiveOnes(nums []int) int {
	var currentCount, maxCount int
	for _, n := range nums {
		if n == 1 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}
	return maxCount
}

