package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var ret [][]uint8
	for i := 0 ; i < dy; i++ {
		s := make([]uint8, dx)
		for j, _ := range(s) {
			// s[j] = uint8((i + j)/2)
			s[j] = uint8(i*j)
		}
		ret = append(ret, s)
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
