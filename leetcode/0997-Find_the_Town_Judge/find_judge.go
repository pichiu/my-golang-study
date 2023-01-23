package test

func findJudge(n int, trust [][]int) int {
	trustPeople := make([]int, n+1)
	trustByOthers := make([]int, n+1)

	for _, item := range trust {
		trustPeople[item[0]]++
		trustByOthers[item[1]]++
	}

	for i := 1; i <= n; i++ {
		if trustPeople[i] == 0 && trustByOthers[i] == n-1 {
			return i
		}
	}
	return -1
}
