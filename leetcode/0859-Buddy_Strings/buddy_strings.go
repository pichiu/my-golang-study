package test

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var diff []int
	m := make(map[byte]int)
    for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
		m[s[i]]++
	}
	ret := false
	if len(diff) == 0 {
		for _, element := range m {
			if element > 1 {
				ret = true
				break
			}
		}
	} else if len(diff) == 2 {
		r := []rune(s)
		r[diff[0]], r[diff[1]] = r[diff[1]], r[diff[0]]
		if string(r) == goal {
			ret = true
		}
	}
	return ret
}
