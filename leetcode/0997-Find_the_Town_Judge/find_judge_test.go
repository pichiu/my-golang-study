package test

import (
	// "fmt"
	"testing"
)

func TestFindJudge(t *testing.T) {
	judgeTests := []struct {
		n           int
		trust       [][]int
		expectedNum int
	}{
		{3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
		{4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, 3},
		{1, [][]int{}, 1},
	}

	for _, v := range judgeTests {
		k := findJudge(v.n, v.trust)

		if k == v.expectedNum {
			t.Log("PASS")
		} else {
			t.Error("Expected=", v.expectedNum, ", wrong answer=", k)
		}
	}
}
