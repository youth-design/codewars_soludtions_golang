package main

import (
	"strings"
	"sync"
)

func reverse(str string) string {
	var reversed strings.Builder
	reversed.Grow(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		reversed.WriteByte(str[i])
	}
	return reversed.String()

}

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	var wg sync.WaitGroup
	for i, word := range words {
		i, word := i, word
		if len(word) >= 5 {
			wg.Add(1)
			go (func() {
				words[i] = reverse(word)
				wg.Done()
			})()
		}
	}

	wg.Wait()

	return strings.Join(words, " ")
}
