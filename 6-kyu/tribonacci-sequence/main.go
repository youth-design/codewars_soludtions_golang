package main

func Tribonacci(signature [3]float64, n int) []float64 {
	res := make([]float64, 0, n)

	for i := 0; i < n; i++ {
		if i < 3 {
			res = append(res, signature[i])
		} else {
			res = append(res, res[len(res)-1]+res[len(res)-2]+res[len(res)-3])
		}
	}
	return res
}
