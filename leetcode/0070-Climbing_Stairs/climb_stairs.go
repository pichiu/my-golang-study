package test

func climbStairs(n int) int {
    l := make([]int, n+1)
	l[0] = 1
	l[1] = 1
	if n > 1 {
		for i := 2; i < n+1; i++ {
			l[i] = l[i-2] + l[i-1]
		}
		return l[n]
	} else {
		return 1
	}
}
