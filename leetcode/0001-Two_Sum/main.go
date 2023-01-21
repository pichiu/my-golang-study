package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
    for i, num1 := range nums {
		res := target - num1
		// fmt.Printf("i=%d, res=%d\n", i, res)
		m[res] = i
	}
	for j, num2 := range nums {
		elem, ok := m[num2]
		// fmt.Printf("elem=%d, ok=%v\n", elem, ok)
		if ok && j != elem {
			return []int{j, elem}
		}
	}
	return []int{-1, -1}
}

func main() {
    fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}
