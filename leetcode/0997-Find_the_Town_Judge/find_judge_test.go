package test

import (
	// "fmt"
	"testing"
)

func TestFindJudge(t *testing.T) {
	// n := 3
	// trust := [][]int{
	// 	{1,3},
	// 	{2,3},
	// 	{3,1},
	// }
	// expectedNum := -1

	// n := 4
	// trust := [][]int{
	// 	{1,3},
	// 	{1,4},
	// 	{2,3},
	// 	{2,4},
	// 	{4,3},
	// }
	// expectedNum := 3

	n := 1
	var trust [][]int
	expectedNum := 1


	k := findJudge(n, trust)
	
	if k == expectedNum {
		t.Log("PASS")
	} else {
		t.Error("Expected=", expectedNum, ", wrong answer=", k)
	}
}
