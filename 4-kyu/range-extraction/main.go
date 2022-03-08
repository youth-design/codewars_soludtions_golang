package main

import (
	"fmt"
	"strings"
)

func Solution(list []int) string {
	res := make([]string, 0)

	interval := make([]int, 0)

	for i, num := range list {
		if len(interval) == 0 || interval[len(interval)-1] == num-1 {
			interval = append(interval, num)
		} else if len(interval) != 0 {
			if len(interval) == 1 {
				res = append(res, fmt.Sprintf("%d", interval[0]))
			} else if len(interval) == 2 {
				res = append(res, fmt.Sprintf("%d", interval[0]))
				res = append(res, fmt.Sprintf("%d", interval[1]))
			} else {
				res = append(res, fmt.Sprintf("%d", interval[0])+"-"+fmt.Sprintf("%d", interval[len(interval)-1]))
			}
			interval = make([]int, 0, 1)
			interval = append(interval, num)
		} else {
			res = append(res, fmt.Sprintf("%d", num))
		}
		if i == len(list)-1 {
			if len(interval) == 1 {
				res = append(res, fmt.Sprintf("%d", interval[0]))
			} else if len(interval) == 2 {
				res = append(res, fmt.Sprintf("%d", interval[0]))
				res = append(res, fmt.Sprintf("%d", interval[1]))
			} else {
				res = append(res, fmt.Sprintf("%d", interval[0])+"-"+fmt.Sprintf("%d", interval[len(interval)-1]))
			}
		}

	}
	return strings.Join(res, ",")
}
