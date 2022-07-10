package main

import (
	"fmt"
)

func HumanReadableTime(seconds int) string {
	res := make([]int, 3)

	res[0] = seconds / 3600
	seconds %= 3600

	res[1] = seconds / 60
	seconds %= 60

	res[2] = seconds

	return fmt.Sprintf("%02d:%02d:%02d", res[0], res[1], res[2])
}
