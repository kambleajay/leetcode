package arrays

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{0, 0, 0}, 0},
		{[]int{0, 1, 0}, 1},
		{[]int{1, 0, 0}, 1},
		{[]int{0, 0, 1}, 1},
		{[]int{1, 0, 1, 0, 1}, 1},
		{[]int{1, 0, 1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 1, 1, 0, 1, 1, 0, 1}, 3},
		{[]int{1, 1, 0, 1, 1, 1, 0, 1}, 3},
		{[]int{}, 0},
	}

	for _, test := range tests {
		if got := findMaxConsecutiveOnes(test.nums); got != test.want {
			t.Errorf("findMaxConsecutiveOnes(%v) = %d", test.nums, got)
		}
	}
}

