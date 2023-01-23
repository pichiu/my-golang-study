package test

func findJudge(n int, trust [][]int) int {
	notJudge := make([]bool, n+1)
	beTrustCount := make([]int, n+1)
	for _, v := range trust {
		notJudge[v[0]] = true
		beTrustCount[v[1]]++
	}
	judge := -1
	m := n - 1
	for i := 1; i <= n; i++ {
		if !notJudge[i] && beTrustCount[i] == m {
			if judge != -1 {
				return -1
			}
			judge = i
		}
	}
	return judge
}
