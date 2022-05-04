package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParticipants(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  int
		Expect int
	}

	cases := []test{
		{
			Name:   "should return 1 for intput 0",
			Value:  0,
			Expect: 1,
		},
		{
			Name:   "should return 2 for input 1",
			Value:  1,
			Expect: 2,
		},
		{
			Name:   "should return 3 for input 3",
			Value:  3,
			Expect: 3,
		},
		{
			Name:   "should return 4 for input 6",
			Value:  6,
			Expect: 4,
		},
		{
			Name:   "should return 5 for input 7",
			Value:  7,
			Expect: 5,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, GetParticipants(c.Value))
		})
	}

}
