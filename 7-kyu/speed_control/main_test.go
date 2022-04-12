package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGps(t *testing.T) {
	t.Parallel()

	type value struct {
		S int
		X []float64
	}

	type test struct {
		Name   string
		Value  value
		Expect int
	}

	cases := []test{
		{
			Name: `should return 41 for input s: 20, x: []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}`,
			Value: value{
				S: 20,
				X: []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61},
			},
			Expect: 41,
		},
		{
			Name: `should return 219 for input s: 12, x: []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}`,
			Value: value{
				S: 12,
				X: []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25},
			},
			Expect: 219,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Gps(c.Value.S, c.Value.X))
		})
	}
}
