package main

import (
	"regexp"
	"strings"
	"unicode"
)

func UpperCaseFirstLetter(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func ToCamelCase(s string) string {
	words := regexp.MustCompile("-|_").Split(s, -1)

	for i, v := range words {
		if i != 0 {
			words[i] = UpperCaseFirstLetter(v)
		}
	}
	return strings.Join(words, "")
}
