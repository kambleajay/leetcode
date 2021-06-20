package arrays

import "testing"

func TestReplaceElems(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
		{[]int{400}, []int{-1}},
		{[]int{}, []int{}},
		{[]int{1,1,1,1,1}, []int{1,1,1,1,-1}},
		{[]int{1,2,3,4,5}, []int{5,5,5,5,-1}},
		{[]int{5,4,3,2,1}, []int{4,3,2,1,-1}},
	}
	for _, test := range tests {
		input := make([]int, len(test.nums))
		copy(input, test.nums)
		got := replaceElements(input)
		if !equal(input, test.want) {
			t.Errorf("replaceElements(%v) = %v (expected %v)", test.nums, input, test.want)
		}
		if !equal(got, test.want) {
			t.Errorf("replaceElements(%v) = %v (expected %v)", test.nums, got, test.want)
		}
	}
}
