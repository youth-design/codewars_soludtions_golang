package main

func MoveZeros(arr []int) []int {
	res := make([]int, len(arr))
	i := 0
	j := len(arr) - 1
	for _, n := range arr {
		if n == 0 {
			res[j] = n
			j -= 1
		} else {
			res[i] = n
			i += 1
		}
	}

	return res
}
