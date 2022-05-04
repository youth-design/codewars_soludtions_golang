package main

import (
	"math"
)

func checkIsPrime(x int) bool {
	if x >= 0 && x < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(x)))
	for i := 2; i <= sqrt; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func Solve(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		if checkIsPrime(i) {
			sum += arr[i]
		}
	}
	return sum
}
