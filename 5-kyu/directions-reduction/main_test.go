package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirReduc(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  []string
		Expect []string
	}

	cases := []test{
		{
			Name:   `["NORTH", "SOUTH", "EAST", "WEST"] should return []`,
			Value:  []string{"NORTH", "SOUTH", "EAST", "WEST"},
			Expect: []string{},
		},
		{
			Name:   `["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"] should return ["WEST"]`,
			Value:  []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"},
			Expect: []string{"WEST"},
		},
		{
			Name:   `["NORTH", "WEST", "SOUTH", "EAST"] should return ["NORTH", "WEST", "SOUTH", "EAST"]`,
			Value:  []string{"NORTH", "WEST", "SOUTH", "EAST"},
			Expect: []string{"NORTH", "WEST", "SOUTH", "EAST"},
		},
		{
			Name:   `["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"] should return ["NORTH"]`,
			Value:  []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"},
			Expect: []string{"NORTH"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, DirReduc(c.Value))
		})
	}
}
