package main

func FindOdd(seq []int) int {
	res := make(map[int]int, len(seq))

	for _, i := range seq {
		res[i] += 1
	}
	for k, v := range res {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}
