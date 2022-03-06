package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidIp(t *testing.T) {
	t.Parallel()

	type test struct {
		Name   string
		Value  string
		Expect bool
	}
	cases := []test{
		{
			Name:   "should test that 12.255.56.1 is correct",
			Value:  "12.255.56.1",
			Expect: true,
		},
		{
			Name:   "should test that '' is uncorrect",
			Value:  "",
			Expect: false,
		},
		{
			Name:   "should test that abc.def.ghi.jkl is uncorrect",
			Value:  "abc.def.ghi.jkl",
			Expect: false,
		},
		{
			Name:   "should test that 123.456.789.0 is uncorrect",
			Value:  "123.456.789.0",
			Expect: false,
		},
		{
			Name:   "should test that 12.34.56 is uncorrect",
			Value:  "12.34.56",
			Expect: false,
		},
		{
			Name:   "should test that 12.34.56 .1 is uncorrect",
			Value:  "12.34.56 .1",
			Expect: false,
		},
		{
			Name:   "should test that 12.34.56.-1 is uncorrect",
			Value:  "12.34.56.-1",
			Expect: false,
		},
		{
			Name:   "should test that 123.045.067.089 is uncorrect",
			Value:  "123.045.067.089",
			Expect: false,
		},
		{
			Name:   "should test that 127.1.1.0 is correct",
			Value:  "127.1.1.0",
			Expect: true,
		},
		{
			Name:   "should test that 0.0.0.0 is correct",
			Value:  "0.0.0.0",
			Expect: true,
		},
		{
			Name:   "should test that 0.34.82.53 is correct",
			Value:  "0.34.82.53",
			Expect: true,
		},
		{
			Name:   "should test that 192.168.1.300 is uncorrect",
			Value:  "192.168.1.300",
			Expect: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Expect, Is_valid_ip(c.Value))
		})
	}
}
