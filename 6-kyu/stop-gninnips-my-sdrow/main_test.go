package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpinWords(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  string
		Expect string
	}

	cases := []test{
		{
			Name:   `should return "emocleW" for input "Welcome"`,
			Value:  "Welcome",
			Expect: "emocleW",
		},
		{
			Name:   `should return "to" for input "to"`,
			Value:  "to",
			Expect: "to",
		},
		{
			Name:   `should return "sraWedoC" for input "CodeWars"`,
			Value:  "CodeWars",
			Expect: "sraWedoC",
		},
		{
			Name:   `should return "Hey wollef sroirraw" for input "Hey fellow warriors"`,
			Value:  "Hey fellow warriors",
			Expect: "Hey wollef sroirraw",
		},
		{
			Name:   `should return "sregruB are my etirovaf tiurf" for input "Burgers are my favorite fruit"`,
			Value:  "Burgers are my favorite fruit",
			Expect: "sregruB are my etirovaf tiurf",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, c.Expect, SpinWords(c.Value))
		})
	}
}
