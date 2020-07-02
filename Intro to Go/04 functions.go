package main

import (
	"fmt"
	"testing"
)

// Create a new function called Hello
// It should accept one string parameter
// It should return a string saying Hello with the value of parameter
// Once you are done, run the code
// Test should PASS

// Hint: you can either use string concatenation or fmt.Sprintf
// Documentation link: https://golang.org/pkg/fmt/#Sprintf

func Hello(str_name string) {
	// fmt.Println("Hello %s", str_name)
	return fmt.Sprintf("Hello, %s!", str_name)
}

func TestHello(t *testing.T) {
	const expected = "Hello, Go!"

	actual := Hello("Go")

	if actual != expected {
		t.Fatalf("expected '%s' but got '%s'", expected, actual)
	}
}
