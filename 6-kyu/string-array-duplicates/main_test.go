package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDub(t *testing.T) {
	t.Parallel()
	type testcase struct {
		Name   string
		Expect []string
		Value  []string
	}
	cases := []testcase{
		{
			Name:   `test ["ccooddddddewwwaaaaarrrrsssss", "piccaninny", "hubbubbubboo"] shold be ["codewars", "picaniny", "hubububo"]`,
			Value:  []string{"ccooddddddewwwaaaaarrrrsssss", "piccaninny", "hubbubbubboo"},
			Expect: []string{"codewars", "picaniny", "hubububo"},
		},
		{
			Name:   `test ["abracadabra", "allottee", "assessee"] should be ["abracadabra", "alote", "asese"]`,
			Value:  []string{"abracadabra", "allottee", "assessee"},
			Expect: []string{"abracadabra", "alote", "asese"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Dup(c.Value))
		})
	}

}
