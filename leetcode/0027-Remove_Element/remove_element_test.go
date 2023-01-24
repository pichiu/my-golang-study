package test

import (
	// "fmt"
	"testing"
	"sort"
)

func TestRemoveElement(t *testing.T) {
	removeTests := []struct{
		nums         []int
		val          int
		expectedNums []int
	}{
		{[]int{0,1,2,2,3,0,4,2}, 2, []int{0,0,1,3,4}},
		{[]int{3,2,2,3}, 3, []int{2,2}},
	}

	for _, e := range removeTests {
		k := removeElement(e.nums, e.val) // Calls your implementation

		if k == len(e.expectedNums) {
			t.Log("PASS")
		} else {
			t.Error("Unexpected ", k)
		}
		sort.Ints(e.nums[:k]); // Sort the first k elements of nums
		// fmt.Println(nums)
		for i, v := range e.expectedNums {
			if e.nums[i] != v {
				t.Error("Unexpected ", i, "-th value: ", e.nums[i], ", ", v)
			}
		}
	}
}
