package test

import "fmt"

func findJudge(n int, trust [][]int) int {
	if n == 1  && len(trust) == 0 {
        return 1
    }
	type void struct{}
	var member void
	townfolks := make(map[int]void)
	candidates := make(map[int]int)
    for _, v := range trust {
		townfolks[v[0]] = member
		candidates[v[1]]++
	}
	fmt.Println(townfolks, candidates)
	ret := -1
	for key, element := range candidates {
		if _, isExist := townfolks[key]; !isExist {
			if element == n - 1 {
				if ret < 0 {
					ret = key
				} else {
					return -1
				}
			}
		}
	}
	return ret
}
