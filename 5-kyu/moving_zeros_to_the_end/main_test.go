package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeros(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  []int
		Expect []int
	}

	cases := []test{
		{
			Name:   "should test that the solution returns the correct value",
			Value:  []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1},
			Expect: []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, MoveZeros(c.Value))
		})
	}

}
