package arrays

import "testing"

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 0, 0, 1, 2, 3}, []int{1, 2, 3, 0, 0, 0}},
		{[]int{1, 2, 3, 0, 0, 0}, []int{1, 2, 3, 0, 0, 0}},
	}
	for _, test := range tests {
		input := make([]int, len(test.nums))
		copy(input, test.nums)
		if moveZeroes(input); !equal(input, test.want) {
			t.Errorf("moveZeros(%v) = %v (expected %v)", test.nums, input, test.want)
		}
	}
}
