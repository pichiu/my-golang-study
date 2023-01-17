package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i, _ := range s {
		s[i] = make([]uint8, dx)
		for j, _ := range(s[i]) {
			// s[i][j] = uint8((i + j)/2)
			s[i][j] = uint8(i^j)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
