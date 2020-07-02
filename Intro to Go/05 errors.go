package main

import "fmt"

func PrintHelloWorld() error {
	_, err := fmt.Println("Hello world")
	if err != nil { // An error occurred
		return err
	}
	return nil
}
