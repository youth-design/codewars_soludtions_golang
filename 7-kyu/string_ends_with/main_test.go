package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  [2]string
		Expect bool
	}

	cases := []test{
		{
			Name:   `should return true for input str: "", ending: ""`,
			Value:  [2]string{"", ""},
			Expect: true,
		},
		{
			Name:   `should return true for input str: " ", ending: ""`,
			Value:  [2]string{" ", ""},
			Expect: true,
		},
		{
			Name:   `should return true for input str: "abc", ending: "c"`,
			Value:  [2]string{"abc", "c"},
			Expect: true,
		},
		{
			Name:   `should return true for input str: "banana", ending: "ana"`,
			Value:  [2]string{"banana", "ana"},
			Expect: true,
		},
		{
			Name:   `should return false for input str: "z", ending: "a"`,
			Value:  [2]string{"z", "a"},
			Expect: false,
		},
		{
			Name:   `should return false for input str: "", ending: "t"`,
			Value:  [2]string{"", "t"},
			Expect: false,
		},
		{
			Name:   `should return false for input str: "$a = $b + 1", ending: "+1"`,
			Value:  [2]string{"$a = $b + 1", "+1"},
			Expect: false,
		},
		{
			Name:   `should return true for input str: "    ", ending: "   "`,
			Value:  [2]string{"    ", "   "},
			Expect: true,
		},
		{
			Name:   `should return false for input str: "1oo", ending: "100"`,
			Value:  [2]string{"1oo", "100"},
			Expect: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, solution(c.Value[0], c.Value[1]))
		})
	}
}
