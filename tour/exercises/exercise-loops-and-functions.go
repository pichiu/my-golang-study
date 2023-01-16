package main

import (
	"fmt"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) float64 {
	t, z := 0., 1.0
	for i := 0; i < 10; i++ {
		t, z = z, z-(z*z-x)/(2*z)
		fmt.Printf("%v: tmp: %g, exec: %g\n", i, t, z)
		if abs(z-t) < 1e-8 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
