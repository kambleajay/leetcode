package arrays

import "testing"

func TestSortByParity(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{3, 1, 2, 4}, []int{4, 2, 1, 3}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{2}},
		{[]int{2,3}, []int{2,3}},
		{[]int{1,2,3,4,5,6,7}, []int{6,2,4,3,5,1,7}},
		{[]int{1,2,3,4,5,6,7,8}, []int{8,2,6,4,5,3,7,1}},
		{[]int{2,2,2}, []int{2,2,2}},
		{[]int{2,4,6,8}, []int{2,4,6,8}},
		{[]int{5,5,5}, []int{5,5,5}},
		{[]int{1,3,5,7}, []int{1,3,5,7}},
		{[]int{0,1,2}, []int{0,2,1}},
	}
	for _, test := range tests {
		input := make([]int, len(test.nums))
		copy(input, test.nums)
		if sortArrayByParity(input); !equal(input, test.want) {
			t.Errorf("sortArrayByParity(%v) = %v (expected %v)", test.nums, input, test.want)
		}
	}
}
