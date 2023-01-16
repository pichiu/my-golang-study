package main

import (
	"fmt"
	"math"
)

const DELTA = 1e-8

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
		if abs(z-t) < DELTA {
			return z
		}
	}
	return z
}

func main() {
	values := []float64{2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range values {
		s, q := Sqrt(v), math.Sqrt(v)
		fmt.Printf("Sqrt: %g, math: %g, diff: %g\n\n", s, q, abs(s-q))
	}
}
