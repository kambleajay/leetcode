package arrays

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 0
	for ; i < len(arr) - 2 && arr[i] < arr[i+1]; i++ {
	}
	if i == len(arr) - 1 || i == 0 {
		return false
	}

	j := len(arr) - 1
	for ; j > 0 && arr[j] < arr[j-1]; j-- {
	}
	if i != j {
		return false
	}
	return true
}
