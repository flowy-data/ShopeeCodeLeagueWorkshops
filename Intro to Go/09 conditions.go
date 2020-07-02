// https://play.golang.org/p/jhz2Fwgx-Xd

package main

import (
	"strconv"
	"testing"
)

// Documentation link: https://golang.org/pkg/strconv/#Itoa

// Return Fizz is multiple of 3
// Return Buzz is multiple of 5
// Otherwise return the number

// if  the number is multiple of 3 and 5, return FizzBuzz

func FizzBuzz(n int) string {

	str_return := ""

	if n%3 == 0 {
		str_return += "Fizz"
	}
	if n%5 == 0 {
		str_return += "Buzz"
	}
	if str_return == "" {
		str_return += strconv.Itoa(n)
	}
	return str_return
}

func TestFizzBuzz(t *testing.T) {
	var testcases = []struct {
		n        int
		expected string
	}{
		{1, "1"},
		{9, "Fizz"},
		{25, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, test := range testcases {
		actual := FizzBuzz(test.n)

		if actual != test.expected {
			t.Fatalf("expected '%s' but got '%s'", test.expected, actual)
		}
	}
}
