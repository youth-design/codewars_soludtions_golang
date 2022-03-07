package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGap(t *testing.T) {
	t.Parallel()

	type testValue struct {
		G int
		M int
		N int
	}

	type test struct {
		Name   string
		Value  testValue
		Expect []int
	}

	cases := []test{
		{
			Name:   `2,100,110 should return [101, 103]`,
			Value:  testValue{G: 2, M: 100, N: 110},
			Expect: []int{101, 103},
		},
		{
			Name:   `4,100,110 should return [103, 107]`,
			Value:  testValue{G: 4, M: 100, N: 110},
			Expect: []int{103, 107},
		},
		{
			Name:   `6,100,110 should return nil`,
			Value:  testValue{G: 6, M: 100, N: 110},
			Expect: nil,
		},
		{
			Name:   `8,300,400 should return [359, 367]`,
			Value:  testValue{G: 8, M: 300, N: 400},
			Expect: []int{359, 367},
		},
		{
			Name:   `10,300,400 should return [337, 347]`,
			Value:  testValue{G: 10, M: 300, N: 400},
			Expect: []int{337, 347},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Gap(c.Value.G, c.Value.M, c.Value.N))
		})
	}
}
