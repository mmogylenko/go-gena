package main

import "math"

// SquareRoot calculation
func SquareRoot() float64 {
	var x float64 = 0.001
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}
