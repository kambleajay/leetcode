package arrays

import (
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	} {
		{[]int{-4,-1,0,3,10}, []int{0,1,9,16,100}},
		{[]int{-7,-3,2,3,11}, []int{4,9,9,49,121}},
		{[]int{-10,-7,-5,-2}, []int{4,25,49,100}},
		{[]int{1,1}, []int{1,1}},
		{[]int{}, []int{}},
		{[]int{-1}, []int{1}},
	}
	for _, test := range tests {
		if got := sortedSquares(test.nums); !equal(got, test.want) {
			t.Errorf("sortedSquares(%v) = %v (expected %v)", test.nums, got, test.want)
		}
	}
}

func benchmarkSortedSquares(b *testing.B, size int) {
	nums := randomIntsBetween(size, -10000, 10000)
	for i := 0; i < b.N; i++ {
		sortedSquares(nums)
	}
}

func BenchmarkSortedSquares10(b *testing.B) { benchmarkSortedSquares(b, 10) }
func BenchmarkSortedSquares1000(b *testing.B) { benchmarkSortedSquares(b, 1000) }
func BenchmarkSortedSquares1000000(b *testing.B) { benchmarkSortedSquares(b, 1000000) }
