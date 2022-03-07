package main

import (
	"math"
)

func checkIsPrime(x int) bool {
	sqrt := int(math.Sqrt(float64(x)))
	for i := 2; i <= sqrt; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func Gap(g, m, n int) []int {
	prevPrime := 0
	for i := m; i <= n; i++ {
		if checkIsPrime(i) {
			if i-prevPrime == g {
				return []int{prevPrime, i}
			}
			prevPrime = i
		}
	}
	return nil
}
