package arrays

//duplicateZeros modifies arr by duplicating each zero.
func duplicateZeros(arr []int) {
	var count, last int
	lastZeroCanDup := true
	for ; len(arr)-count > 0; last++ {
		if arr[last] != 0 {
			count++
		} else {
			if len(arr)-count > 1 {
				count += 2
			} else {
				count++
				lastZeroCanDup = false
			}
		}
	}
	for dst, src := len(arr)-1, last-1; src >= 0; src-- {
		if arr[src] != 0 {
			arr[dst] = arr[src]
			dst--
		} else {
			if !lastZeroCanDup && src == last - 1 {
				arr[dst] = 0
				dst--
			} else {
				arr[dst], arr[dst-1] = 0, 0
				dst -= 2
			}
		}
	}
}
