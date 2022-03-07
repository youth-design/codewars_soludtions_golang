package main

type Result struct {
	C rune // character
	L int  // count
}

func LongestRepetition(text string) Result {
	rCount := make(map[rune]int)
	result := Result{}

	var tmpR rune
	var c int
	for i, r := range text {
		if tmpR == r {
			c += 1
		} else {
			if rCount[tmpR] < c {
				rCount[tmpR] = c
			}
			tmpR = r

			c = 1
		}
		if i == len(text)-1 && rCount[tmpR] < c {
			rCount[tmpR] = c
		}
		if result.L < c {
			result.C = r
			result.L = c
		}
	}

	return result
}
