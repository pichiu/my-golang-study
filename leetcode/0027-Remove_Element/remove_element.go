package test

// import "fmt"

func removeElement(nums []int, val int) int {
	p := 0
	for _, v := range nums {
		if v != val {
			nums[p] = v
			p++
		}
	}
	// fmt.Println(nums)
	return p
}
