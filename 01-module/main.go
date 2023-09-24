package main

import (
	"fmt"

	"github.com/akash-mohanto/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	// same as above

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
