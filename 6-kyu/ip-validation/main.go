package main

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	blocks := strings.Split(ip, ".")
	if len(blocks) != 4 {
		return false
	}

	for _, b := range blocks {
		i, err := strconv.Atoi(b)
		if err != nil || i < 0 || i > 255 || (b[0] == '0' && len(b) != 1) {
			return false
		}
	}
	return true
}
