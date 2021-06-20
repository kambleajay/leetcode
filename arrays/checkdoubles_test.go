package arrays

import "testing"

func TestCheckDouble(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{10, 2, 5, 3}, true},
		{[]int{7, 1, 14, 11}, true},
		{[]int{3, 1, 7, 11}, false},
		{[]int{}, false},
		{[]int{5}, false},
		{[]int{5,11}, false},
		{[]int{5,0,1,2,3,4,5,6,7,8,9,10}, true},
	}
	for _, test := range tests {
		if got := checkIfExist(test.nums); got != test.want {
			t.Errorf("checkIfExist(%v) = %v (expected %v)", test.nums, got, test.want)
		}
	}
}