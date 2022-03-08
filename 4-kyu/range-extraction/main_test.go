package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type test struct {
		Name   string
		Value  []int
		Expect string
	}

	cases := []test{
		{
			Name:   `[-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20] should return "-6,-3-1,3-5,7-11,14,15,17-20"`,
			Value:  []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20},
			Expect: "-6,-3-1,3-5,7-11,14,15,17-20",
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
