package main

import "fmt"

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int, len(nums))

    for i, num := range nums { 
        rem := target - num
        if j, found := visited[rem]; found {
            return []int{i, j}
        }
        visited[num] = i
    }
    return nil
}

func main() {
    fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}
