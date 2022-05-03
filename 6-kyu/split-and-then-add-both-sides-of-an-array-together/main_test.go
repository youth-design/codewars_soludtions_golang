package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitAndAdd(t *testing.T) {
	t.Parallel()

	type value struct {
		Numbers []int
		n       int
	}

	type test struct {
		Name   string
		Value  value
		Expect []int
	}

	cases := []test{
		{
			Name: "should return []int{5, 10} for input numbers: []int{1, 2, 3, 4, 5}, n: 2",
			Value: value{
				Numbers: []int{1, 2, 3, 4, 5},
				n:       2,
			},
			Expect: []int{5, 10},
		},
		{
			Name: "should return []int{15} for input numbers: []int{1, 2, 3, 4, 5}, n: 3",
			Value: value{
				Numbers: []int{1, 2, 3, 4, 5},
				n:       3,
			},
			Expect: []int{15},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, SplitAndAdd(c.Value.Numbers, c.Value.n))
		})
	}
}
