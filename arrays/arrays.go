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

//digitCount returns the number of digits in n.
func digitCount(n int) int {
	count := 0
	for ;n != 0; count++ {
		n = n / 10
	}
	return count
}

//findNumbers returns the number of two-digit numbers in nums.
func findNumbers(nums []int) int {
	countOfEvenNumOfDigits := 0
	for _, n := range nums {
		if (digitCount(n) % 2) == 0 {
			countOfEvenNumOfDigits++
		}
	}
	return countOfEvenNumOfDigits
}
