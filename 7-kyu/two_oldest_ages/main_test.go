package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoOldestAges(t *testing.T) {
	t.Parallel()
	type test struct {
		Name   string
		Value  []int
		Expect [2]int
	}

	cases := []test{
		{
			Name:   "should return 18 and 83 for input []int{6,5,83,5,3,18}",
			Value:  []int{6, 5, 83, 5, 3, 18},
			Expect: [2]int{18, 83},
		},
		{
			Name:   "should return 45 and 87 for input []int{1,5,87,45,8,8}",
			Value:  []int{1, 5, 87, 45, 8, 8},
			Expect: [2]int{45, 87},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, TwoOldestAges(c.Value))
		})
	}

}
