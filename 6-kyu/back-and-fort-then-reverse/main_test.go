package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrange(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  []int
		Expect []int
	}

	cases := []test{
		{
			Name:   "should return []int{1, 2} for input []int{1, 2}",
			Expect: []int{1, 2},
			Value:  []int{1, 2},
		},
		{
			Name:   "should return []int{4, 3, 2} for input []int{4, 2, 3}",
			Expect: []int{4, 3, 2},
			Value:  []int{4, 2, 3},
		},
		{
			Name:   "should return []int{9, 1, 5, 7, -2, 6, -3, 8, 5} for input []int{9, 7, -2, 8, 5, -3, 6, 5, 1}",
			Expect: []int{9, 1, 5, 7, -2, 6, -3, 8, 5},
			Value:  []int{9, 7, -2, 8, 5, -3, 6, 5, 1},
		},
		{
			Name:   "should return nil for input nil",
			Expect: nil,
			Value:  nil,
		},
		{
			Name:   "should return []int{1} for input []int{1}",
			Expect: []int{1},
			Value:  []int{1},
		},
		{
			Name:   "should return []int{2, 4, 3, 4} for input []int{2, 4, 3, 4}",
			Expect: []int{2, 4, 3, 4},
			Value:  []int{2, 4, 3, 4},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Arrange(c.Value))
		})

	}
}
