package arrays

import "testing"

func TestMountainArr(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
		{[]int{0,2,3,4,5,2,1,0}, true},
		{[]int{0,2,3,3,5,2,1,0}, false},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{0,1,2,3,4,5,6,7,8,9}, false},
		{[]int{9,8,7,6,5,4,3,2,1,0}, false},
	}
	for _, test := range tests {
		if got := validMountainArray(test.nums); got != test.want {
			t.Errorf("validMountainArray(%v) = %v (expected %v)", test.nums, got, test.want)
		}
	}
}