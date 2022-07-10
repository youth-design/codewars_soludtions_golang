package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumanReadableTime(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Expect string
		Value  int
	}

	cases := []test{
		{
			Name:   `should return "00:00:00" for input 0`,
			Expect: "00:00:00",
			Value:  0,
		},
		{
			Name:   `should return "00:00:00" for input 59`,
			Expect: "00:00:59",
			Value:  59,
		},
		{
			Name:   `should return "00:01:00" for input 60`,
			Expect: "00:01:00",
			Value:  60,
		},
		{
			Name:   `should return "00:01:30" for input 90`,
			Expect: "00:01:30",
			Value:  90,
		},
		{
			Name:   `should return "00:59:59" for input 3599`,
			Expect: "00:59:59",
			Value:  3599,
		},
		{
			Name:   `should return "01:00:00" for input 3600`,
			Expect: "01:00:00",
			Value:  3600,
		},
		{
			Name:   `should return "12:34:56" for input 45296`,
			Expect: "12:34:56",
			Value:  45296,
		},
		{
			Name:   `should return "23:59:59" for input 86399`,
			Expect: "23:59:59",
			Value:  86399,
		},
		{
			Name:   `should return "24:00:00" for input 86400`,
			Expect: "24:00:00",
			Value:  86400,
		},
		{
			Name:   `should return "99:59:59" for input 359999`,
			Expect: "99:59:59",
			Value:  359999,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, HumanReadableTime(c.Value))
		})
	}
}
