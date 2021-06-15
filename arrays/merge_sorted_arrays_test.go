package arrays

import "testing"

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{1, 1, 0, 0}, 2, []int{1, 1}, 2, []int{1, 1, 1, 1}},
		{[]int{1, 2, 0, 0}, 2, []int{1, 2}, 2, []int{1, 1, 2, 2}},
		{[]int{}, 0, []int{}, 0, []int{}},
		{[]int{10,20,30,0,0,0}, 3, []int{40,50,60}, 3, []int{10,20,30,40,50,60}},
	}
	for _, test := range tests {
		input := make([]int, len(test.nums1))
		copy(input, test.nums1)
		if merge(input, test.m, test.nums2, test.n); !equal(input, test.want) {
			t.Errorf("merge(%v) = %v (expected %v)", test.nums1, input, test.want)
		}
	}
}

func benchmarkMerge(b *testing.B, size int) {
	nums1 := randomIntsBetween(size, -1000000000, 1000000000)
	initialLen := len(nums1)
	nums2 := randomIntsBetween(size, -1000000000, 1000000000)
	zeroArr := make([]int, len(nums1))
	nums1 = append(nums1, zeroArr...)
	for i := 0; i < b.N; i++ {
		merge(nums1, initialLen, nums2, len(nums2))
	}
}

func BenchmarkMerge10(b *testing.B)      { benchmarkMerge(b, 10) }
func BenchmarkMerge1000(b *testing.B)    { benchmarkMerge(b, 1000) }
func BenchmarkMerge1000000(b *testing.B) { benchmarkMerge(b, 1000000) }
