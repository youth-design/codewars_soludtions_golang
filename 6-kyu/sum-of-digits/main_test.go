package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitalRoot(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  int
		Expect int
	}

	cases := []test{
		{
			Name:   "should return 7 for input 16",
			Value:  16,
			Expect: 7,
		},
		{
			Name:   "should return 6 for input 195",
			Value:  195,
			Expect: 6,
		},
		{
			Name:   "should return 2 for input 992",
			Value:  992,
			Expect: 2,
		},
		{
			Name:   "should return 9 for input 167346",
			Value:  167346,
			Expect: 9,
		},
		{
			Name:   "should return 0 for input 0",
			Value:  0,
			Expect: 0,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, c.Expect, DigitalRoot(c.Value))
		})
	}
}
