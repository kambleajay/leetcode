package arrays

import (
	"testing"
)

func TestDuplicateZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0}, []int{0}},
		{[]int{0, 1}, []int{0, 0}},
		{[]int{1, 0, 0}, []int{1, 0, 0}},
		{[]int{0, 1, 0}, []int{0, 0, 1}},
		{[]int{0, 0, 1}, []int{0, 0, 0}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{1}, []int{1}},
		{[]int{8, 4, 5, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0}},
		{[]int{0, 4, 1, 0, 0, 8, 0, 0, 3}, []int{0, 0, 4, 1, 0, 0, 0, 0, 8}},
		{[]int{1, 5, 2, 0, 6, 8, 0, 6, 0}, []int{1, 5, 2, 0, 0, 6, 8, 0, 0}},
	}
	for _, test := range tests {
		input := make([]int, len(test.nums))
		copy(input, test.nums)
		if duplicateZeros(input); !equal(input, test.want) {
			t.Errorf("duplicateZeros(%v) = %v (expected %v)", test.nums, input, test.want)
		}
	}
}

func benchmarkDuplicateZeroes(b *testing.B, size int) {
	nums := randomIntsBetween(size, 0, 9)
	for i := 0; i < b.N; i++ {
		duplicateZeros(nums)
	}
}

func BenchmarkDuplicateZeroes10(b *testing.B)      { benchmarkDuplicateZeroes(b, 10) }
func BenchmarkDuplicateZeroes1000(b *testing.B)    { benchmarkDuplicateZeroes(b, 1000) }
func BenchmarkDuplicateZeroes1000000(b *testing.B) { benchmarkDuplicateZeroes(b, 1000000) }
