package main

func SplitAndAdd(numbers []int, n int) []int {
	if n == 0 {
		return numbers
	}
	if len(numbers)%2 == 0 {
		res := make([]int, 0, len(numbers)/2)
		for i, j := 0, len(numbers)/2; i < len(numbers)/2 && j < len(numbers); i, j = i+1, j+1 {
			res = append(res, numbers[i]+numbers[j])
		}
		return SplitAndAdd(res, n-1)
	} else {
		res := make([]int, 0, len(numbers)/2+1)
		res = append(res, numbers[len(numbers)/2])
		for i, j := 0, len(numbers)/2+1; i < len(numbers)/2 && j < len(numbers); i, j = i+1, j+1 {
			res = append(res, numbers[i]+numbers[j])
		}
		return SplitAndAdd(res, n-1)
	}
}
