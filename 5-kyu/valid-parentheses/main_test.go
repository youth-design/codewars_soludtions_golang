package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  string
		Expect bool
	}

	cases := []test{
		{
			Name:   `expected true for input "()"`,
			Value:  "()",
			Expect: true,
		},
		{
			Name:   `expected false for input ")"`,
			Value:  ")",
			Expect: false,
		},
		{
			Name:   `expected false for input ")(()))"`,
			Value:  ")(()))",
			Expect: false,
		},
		{
			Name:   `expected true for input "(())((()())())`,
			Value:  "(())((()())())",
			Expect: true,
		},
		{
			Name:   `expected false for input "())"`,
			Value:  "())",
			Expect: false,
		},
		{
			Name:   `expected false for input "())("`,
			Value:  "())(",
			Expect: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, c.Expect, ValidParentheses(c.Value))
		})
	}
}
