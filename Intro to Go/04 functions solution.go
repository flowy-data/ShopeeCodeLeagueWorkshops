// works in go playground
package main

import (
	"fmt"
	"testing"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func TestHello(t *testing.T) {
	const expected = "Hello, Go!"

	actual := Hello("Go")

	if actual != expected {
		t.Fatalf("expected %s but got %s", expected, actual)
	}
}
