package arrays

func checkIfExist(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	seen := make(map[int]bool)
	for _, n := range arr {
		if seen[n*2] || (n % 2 == 0 && seen[n/2]) {
			return true
		}
		seen[n] = true
	}
	return false
}