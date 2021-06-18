package arrays

import "testing"

func TestRemoveDups(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{}, 0, []int{}},
		{[]int{0, 0, 0, 1, 1, 2, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1, 1}, 1, []int{1}},
		{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}, 5, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		input := make([]int, len(test.nums))
		copy(input, test.nums)
		got := removeDuplicates(input)
		if got != test.k {
			t.Errorf("removeDuplicates(%v) = %v (expected %v)", test.nums, got, test.k)
		}
		if !equal(input[:test.k], test.want) {
			t.Errorf("removeDuplicates(%v) = %v (expected %v)", test.nums, input, test.want)
		}
	}
}
