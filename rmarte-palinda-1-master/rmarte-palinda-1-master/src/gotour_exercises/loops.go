package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const delta = 0.0000000001
	z := x
	oldZ := 0.0
	numIter := 0
	for math.Abs(oldZ-z) > delta {
		oldZ = z
		z -= (z*z - x) / (2 * z)
		numIter++
	}
	//println("Number of iterations: ", numIter)
	return z
}

func main() {
	for i := 0; i <= 100; i++ {
		n := float64(i)
		myVal := Sqrt(n)
		mathVal := math.Sqrt(n)
		fmt.Println("Value:", n, "My function:", myVal, "Math library:", mathVal, "Difference:", math.Abs(myVal-mathVal))
	}
}
