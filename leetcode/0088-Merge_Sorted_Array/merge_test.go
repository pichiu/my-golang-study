package test

import (
	"testing"
)

func TestMerge(t *testing.T) {
	mergeTests := []struct {
		nums1        []int
		nums2        []int
		m            int
		n            int
		expectedNums []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, []int{}, 1, 0, []int{1}},
		{[]int{0}, []int{1}, 0, 1, []int{1}},
		{[]int{1, 2, 3, 0, 0, 0}, []int{4, 5, 6}, 3, 3, []int{1, 2, 3, 4, 5, 6}},
		{[]int{4, 5, 6, 0, 0, 0}, []int{1, 2, 3}, 3, 3, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, e := range mergeTests {
		merge(e.nums1, e.m, e.nums2, e.n)
		for i, v := range e.expectedNums {
			if e.nums1[i] != v {
				t.Error("Unexpected ", i, "-th value: ", e.nums1[i], ", ", v)
			}
		}
	}
}
