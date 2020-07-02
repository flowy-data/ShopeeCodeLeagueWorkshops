//https://play.golang.org/p/CJeEFOJEFeG
package main

import (
	"strconv"
	"testing"
)

// Documentation link: https://golang.org/pkg/strconv/

func Sum(a, b string) int {
	int_a, _ := strconv.Atoi(a)
	int_b, _ := strconv.Atoi(b)
	return int_a + int_b
}

func TestSum(t *testing.T) {
	const expected = 42

	actual := Sum("40", "2")

	if actual != expected {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
