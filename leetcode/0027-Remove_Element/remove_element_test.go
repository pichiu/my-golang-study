package test

import (
	"fmt"
	"testing"
	"sort"
)

func TestRemoveElement(t *testing.T) {
	// nums := []int{3,2,2,3}
	// val := 3
	// expectedNums := []int{2,2}
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	expectedNums := []int{0,0,1,3,4}
	k := removeElement(nums, val) // Calls your implementation

	if k == len(expectedNums) {
		t.Log("PASS")
	} else {
		t.Error("Unexpected ", k)
	}
	sort.Ints(nums[:k]); // Sort the first k elements of nums
	// fmt.Println(nums)
	for i, v := range expectedNums {
		if nums[i] != v {
			t.Error("Unexpected ", i, "-th value: ", nums[i], ", ", v)
		}
	}
}
