package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOdd(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  []int
		Expect int
	}

	cases := []test{
		{
			Name:   "should return 5 for input []int{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5}",
			Value:  []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
			Expect: 5,
		},
		{
			Name:   "should return -1 for input []int{1,1,2,-2,5,2,4,4,-1,-2,5}",
			Value:  []int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
			Expect: -1,
		},
		{
			Name:   "should return 5 for input []int{20,1,1,2,2,3,3,5,5,4,20,4,5}",
			Value:  []int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
			Expect: 5,
		},
		{
			Name:   "should return 10 for input []int{10}",
			Value:  []int{10},
			Expect: 10,
		},
		{
			Name:   "should return 10 for input []int{1,1,1,1,1,1,10,1,1,1,1}",
			Value:  []int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
			Expect: 10,
		},
		{
			Name:   "should return 1 for input []int{5,4,3,2,1,5,4,3,2,10,10}",
			Value:  []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
			Expect: 1,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, FindOdd(c.Value))
		})
	}
}
