package main

func TwoOldestAges(ages []int) [2]int {
	max2 := -1
	max1 := -1

	for _, i := range ages {
		if i > max2 {
			max1 = max2
			max2 = i
		} else if i > max1 {
			max1 = i
		} else if i == max2 {
			max1 = max2
		}
	}

	return [2]int{max1, max2}
}
