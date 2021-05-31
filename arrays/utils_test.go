package arrays

import (
	"math/rand"
	"time"
)

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

func randomIntBetween(min int, max int) int {
	return gen.Intn(max - min) + min
}

func randomIntsBetween(size int, min int, max int) []int {
	var nums []int
	for i := 0; i < size; i++ {
		nums = append(nums, randomIntBetween(min, max))
	}
	return nums
}

func equal(xs, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}
