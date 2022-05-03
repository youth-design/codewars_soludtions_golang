package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  []int
		Expect int
	}

	cases := []test{
		{
			Name:   "should return 0 for input []int{}",
			Value:  []int{},
			Expect: 0,
		},
		{
			Name:   "should return 7 for input []int{1,2,3,4}",
			Value:  []int{1, 2, 3, 4},
			Expect: 7,
		},
		{
			Name:   "should return 47 for input []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}",
			Value:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Expect: 47,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Solve(c.Value))
		})
	}

}
