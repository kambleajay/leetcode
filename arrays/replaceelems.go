package arrays

func replaceElements(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	last := len(arr) - 1
	i, max := last-1, arr[last]
	arr[last] = -1
	var tmp int
	for ; i > -1; i-- {
		arr[i], tmp = max, arr[i]
		if tmp > max {
			max = tmp
		}
	}
	return arr
}
