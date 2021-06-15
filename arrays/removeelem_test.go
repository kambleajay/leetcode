package arrays

import "testing"

func TestRemoveElem(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		k    int
		want []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2, 3, 3}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 4, 0, 3, 2, 2, 2}},
		{[]int{1, 1, 1, 1}, 1, 0, []int{1, 1, 1, 1}},
		{[]int{1, 1, 1, 1}, 0, 4, []int{1, 1, 1, 1}},
		{[]int{0, 1, 0, 1, 0, 1}, 1, 3, []int{0, 0, 0, 1, 1, 1}},
		{[]int{0, 1, 0, 1, 0, 1, 0}, 1, 4, []int{0, 0, 0, 0, 1, 1, 1}},
		{[]int{0, 1, 0, 1, 0, 1}, 0, 3, []int{1, 1, 1, 0, 0, 0}},
		{[]int{}, 0, 0, []int{}},
	}
	for _, test := range tests {
		input := make([]int, len(test.nums))
		copy(input, test.nums)
		got := removeElement(input, test.val)
		if got != test.k {
			t.Errorf("removeElement(%v) = %v (expected %v)", test.nums, got, test.k)
		}
		if !equal(input, test.want) {
			t.Errorf("removeElement(%v) = %v (expected %v)", test.nums, input, test.want)
		}
	}
}
