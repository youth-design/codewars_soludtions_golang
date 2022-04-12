package main

func reduceNumber(n, count int) int {
	for n > 9 {
		count += n % 10
		n = (n - n%10) / 10
	}
	count += n
	if count > 9 {
		return reduceNumber(count, 0)
	}
	return count
}

func DigitalRoot(n int) int {
	return reduceNumber(n, 0)
}
