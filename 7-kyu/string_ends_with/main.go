package main

func solution(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	return str[len(str)-len(ending):] == ending
}
