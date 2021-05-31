package arrays

//posAndNegSquares squares numbers in nums and returns two slices
//for squares of positive and negative numbers.
func posAndNegSquares(nums []int) ([]int, []int) {
	var posSquares []int
	var negSquares []int
	for _, n := range nums {
		s := n * n
		if n < 0 {
			negSquares = append(negSquares, s)
		} else {
			posSquares = append(posSquares, s)
		}
	}
	return negSquares, posSquares
}

//reverse reverses a slice of ints in place
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

//mergeSquares takes two sorted slices of ints and
//returns a merged slice that is sorted,
func mergeSquares(negSquares []int, posSquares []int) []int {
	var result []int
	var i, j int
	for ; i < len(posSquares) && j < len(negSquares); {
		ps := posSquares[i]
		ns := negSquares[j]
		switch {
		case ps <= ns:
			result = append(result, ps)
			i++
		case ps > ns:
			result = append(result, ns)
			j++
		}
	}
	if i < len(posSquares) {
		result = append(result, posSquares[i:]...)
	} else if j < len(negSquares) {
		result = append(result, negSquares[j:]...)
	}
	return result
}

//sortedSquares takes a sorted slice of ints,
//and returns a sorted slice of squares.
func sortedSquares(nums []int) []int {
	negSquares, posSquares := posAndNegSquares(nums)
	reverse(negSquares)
	return mergeSquares(negSquares, posSquares)
}
