// 03 constants
// this code will fail as constants are only defined once 

package main

import (
	"fmt"
)

func main() {

	//var world string
	// hello = "hello"
	// world = "world"
	const hello, world string
	hello, world = "hello", "world"

	//same results for the 2 lines
	fmt.Println(hello + " " + world)
	fmt.Println(hello, world)


	
}
