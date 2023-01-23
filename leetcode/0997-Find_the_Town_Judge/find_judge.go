package test

func findJudge(n int, trust [][]int) int {
	if n == 1 { return 1 }
    arr := make([]int, n+1)
    for _, tru := range(trust) {
        arr[tru[0]] -= 1
        arr[tru[1]] += 1
    }
    for i, a := range(arr) {
        if a == n - 1 {
            return i
        }
    }
    return -1
}
