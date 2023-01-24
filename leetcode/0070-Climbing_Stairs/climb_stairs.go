package test

func climbStairs(n int) int {
	current, next := 1, 1
	for i := 1; i < n; i++ {
		current, next = next, next+current
	}
	return next
}
