package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  string
		Expect string
	}

	cases := []test{
		{
			Name:   `"" should return ""`,
			Value:  "",
			Expect: "",
		},
		{
			Name:   `"The_Stealth_Warrior" should return "TheStealthWarrior"`,
			Value:  "The_Stealth_Warrior",
			Expect: "TheStealthWarrior",
		},
		{
			Name:   `"the-Stealth-Warrior" should return "theStealthWarrior"`,
			Value:  "the-Stealth-Warrior",
			Expect: "theStealthWarrior",
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, ToCamelCase(c.Value))
		})
	}
}
