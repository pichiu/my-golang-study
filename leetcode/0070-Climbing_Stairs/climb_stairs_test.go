package test

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	climbTests := []struct {
		n           int
		expectedNum int
	}{
		{2, 2},
		{3, 3},
		{5, 8},
		{1, 1},
	}

	for _, e := range climbTests {
		k := climbStairs(e.n)

		if k == e.expectedNum {
			t.Log("PASS")
		} else {
			t.Error("Unexpected ", k)
		}
	}
}