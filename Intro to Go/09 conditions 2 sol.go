package main

import (
	"strconv"
	"strings"
	"testing"
)

func AllFizzBuzz(n string) string {
	max, _ := strconv.Atoi(n)

	fizzbuzz := []string{}

	for i := 1; i <= max; i++ {
		fizzbuzz = append(fizzbuzz, FizzBuzz(i))
	}

	return strings.Join(fizzbuzz, " ")
}

func FizzBuzz(n int) string {
	fizzbuzz := ""

	if n%3 == 0 {
		fizzbuzz = "Fizz"
	}
	if n%5 == 0 {
		fizzbuzz += "Buzz"
	}

	if fizzbuzz == "" {
		return strconv.Itoa(n)
	}

	return fizzbuzz
}

func TestAllFizzBuzz(t *testing.T) {
	const expected = "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz"

	actual := AllFizzBuzz("15")

	if actual != expected {
		t.Fatalf("expected '%s' but got '%s'", expected, actual)
	}
}
