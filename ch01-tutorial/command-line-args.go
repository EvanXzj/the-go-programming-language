// prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// :=  can only within a function
	length := len(os.Args)

	for i := 0; i < length; i++ {
		// The first element of os.Args, os.Args[0] is the name of the command itself.
		s += sep + os.Args[i]

		sep = " "
	}

	fmt.Println(s)
}
