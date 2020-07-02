// https://play.golang.org/p/u1Qz9VthBLa
package main

import (
	"testing"
)

func IsEven(n int) bool {
	return n%2 == 0
}

func TestIsEven(t *testing.T) {
	const expected = true

	actual := IsEven(2)

	if actual != expected {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}
