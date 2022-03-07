package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestRepetition(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  string
		Expect Result
	}

	cases := []test{
		{
			Name:   `"aaaabb" should return Result{'a', 4}`,
			Value:  "aaaabb",
			Expect: Result{C: 'a', L: 4},
		},
		{
			Name:   `"bbbaaabaaaa" should return Result{'a', 4}`,
			Value:  "bbbaaabaaaa",
			Expect: Result{C: 'a', L: 4},
		},
		{
			Name:   `"cbdeuuu900" should return Result{'a', 4}`,
			Value:  "cbdeuuu900",
			Expect: Result{C: 'u', L: 3},
		},
		{
			Name:   `"abbbbb" should return Result{'a', 4}`,
			Value:  "abbbbb",
			Expect: Result{C: 'b', L: 5},
		},
		{
			Name:   `"" should return Result{}`,
			Value:  "",
			Expect: Result{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, c.Expect, LongestRepetition(c.Value))
		})
	}

}
