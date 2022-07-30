package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  []int
		Expect int
	}

	cases := []test{
		{
			Name:   "Expect 9 for input []int{9}",
			Expect: 9,
			Value:  []int{9},
		},
		{
			Name:   "Expect 9 for input []int{6,9,21}",
			Expect: 9,
			Value:  []int{6, 9, 21},
		},
		{
			Name:   "Expect 9 for input []int{1,21,55}",
			Expect: 3,
			Value:  []int{1, 21, 55},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Solution(c.Value))
		})
	}
}
