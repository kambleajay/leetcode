package arrays

//digitCount returns the number of digits in n.
func digitCount(n int) int {
	count := 0
	for ; n != 0; count++ {
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

