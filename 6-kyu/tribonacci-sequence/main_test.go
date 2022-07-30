package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTribonacci(t *testing.T) {
	t.Parallel()

	type test struct {
		Name           string
		ValueSignature [3]float64
		ValueN         int
		Expect         []float64
	}

	cases := []test{
		{
			Name:           `Expect []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105} for input signature: [3]float64{1, 1, 1}, n: 10`,
			ValueN:         10,
			ValueSignature: [3]float64{1, 1, 1},
			Expect:         []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105},
		},
		{
			Name:           `Expect []float64{1, 2, 3, 6, 11, 20, 37, 68, 125, 230} for input signature: [3]float64{1, 2, 3}, n: 10`,
			ValueN:         10,
			ValueSignature: [3]float64{1, 2, 3},
			Expect:         []float64{1, 2, 3, 6, 11, 20, 37, 68, 125, 230},
		},
		{
			Name:           `Expect []float64{} for input signature: [3]float64{1, 2, 3}, n: 0`,
			ValueN:         0,
			ValueSignature: [3]float64{300, 200, 100},
			Expect:         []float64{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Tribonacci(c.ValueSignature, c.ValueN))
		})
	}
}
