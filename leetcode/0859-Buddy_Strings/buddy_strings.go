package test

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		m := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			if m[s[i]] == 1 {
				return true
			}
			m[s[i]]++
		}
		return false
	}
	var diff []int
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}
	if len(diff) == 2 {
		r := []rune(s)
		g := []rune(goal)
		return r[diff[0]] == g[diff[1]] && r[diff[1]] == g[diff[0]]
	}
	return false
}
