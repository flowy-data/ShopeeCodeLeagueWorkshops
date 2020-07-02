// strings

package main

import (
	"reflect"
	"strings"
	"testing"
)

// Hint: Split the string into chunks
// Documentation link: https://golang.org/pkg/strings

func GetNumbers(in string) []string {
	return strings.Split(in, ", ")
}

func TestGetNumbers(t *testing.T) {
	const given = "0, 1, 1, 2, 3, 5, 8, 13, 21, 34"
	expected := []string{"0", "1", "1", "2", "3", "5", "8", "13", "21", "34"}

	actual := GetNumbers(given)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}
