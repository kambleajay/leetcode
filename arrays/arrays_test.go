package arrays

import (
	"math/rand"
	"testing"
	"time"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1,1,0,1,1,1}, 3},
		{[]int{0,0,0}, 0},
		{[]int{0,1,0},1},
		{[]int{1,0,0}, 1},
		{[]int{0,0,1}, 1},
		{[]int{1,0,1,0,1}, 1},
		{[]int{1,0,1,1,0,1,1,1}, 3},
		{[]int{1,1,1,0,1,1,0,1}, 3},
		{[]int{1,1,0,1,1,1,0,1}, 3},
		{[]int{}, 0},
	}

	for _, test := range tests {
		if got := findMaxConsecutiveOnes(test.nums); got != test.want {
			t.Errorf("findMaxConsecutiveOnes(%v) = %d", test.nums, got)
		}
	}
}

func TestDigitCount(t *testing.T) {
	var tests = []struct {
		n int
		want int
	}{
		{12, 2},
		{1, 1},
		{123, 3},
		{-1, 1},
	}

	for _, test := range tests {
		if got := digitCount(test.n); got != test.want {
			t.Errorf("digitCount(%d) = %d", test.n, got)
		}
	}
}

func TestFindNumbers(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{12,345,2,6,7896}, 2},
		{[]int{555,901,482,1771}, 1},
		{[]int{},0},
		{[]int{12,23,34}, 3},
		{[]int{123,456,789,10}, 1},
	}

	for _, test := range tests {
		if got := findNumbers(test.nums); got != test.want {
			t.Errorf("findNumbers(%v) = %d", test.nums, got)
		}
	}
}

func randomGen() *rand.Rand {
	seed := time.Now().UTC().UnixNano()
	return rand.New(rand.NewSource(seed))
}

var gen = randomGen()

func randomInt(limit int) int {
	return gen.Intn(limit)
}

func randomInts(size int, limit int) []int {
	var nums []int
	for i := 0; i < size; i++ {
		nums = append(nums, randomInt(limit))
	}
	return nums
}

func benchmarkFindNumbers(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		findNumbers(randomInts(size, 1000000))
	}
}

func BenchmarkFindNumbers10(b *testing.B)   { benchmarkFindNumbers(b, 10) }
func BenchmarkFindNumbers1000(b *testing.B)   { benchmarkFindNumbers(b, 1000) }
func BenchmarkFindNumbers1000000(b *testing.B)   { benchmarkFindNumbers(b, 1000000) }
