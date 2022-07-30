package main

import (
	"fmt"
)

func GCD(a, b int) int {
	c := 0
	for b > 0 {
		c = a % b
		a = b
		b = c
	}

	if a < 0 {
		a *= -1
	}
	return a
}

func Solution(ar []int) int {
	min := ar[0]
	max := ar[0]

	if len(ar) == 1 {
		return ar[0]
	}

	for _, n := range ar {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	gcd := GCD(max, min)

	for _, n := range ar {
		if n%gcd != 0 {
			gcd = GCD(n, min)
		}
	}

	return gcd * len(ar)
}

func main() {
	data := []int{6, 9, 21, 24, 27}
	fmt.Println(data, Solution(data))
}
