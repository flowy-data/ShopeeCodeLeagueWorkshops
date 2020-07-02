// Given a number of iteration, print all the values of FizzBuzz
// not complete
package main

import (
	"reflect"
	"testing"
)

func GetNumbers(in string) []string {
	return []string{}
}

func TestGetNumbers(t *testing.T) {
	const given = "0, 1, 1, 2, 3, 5, 8, 13, 21, 34"
	expected := []string{"0", "1", "1", "2", "3", "5", "8", "13", "21", "34"}

	actual := GetNumbers(given)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}
