package main

import (
	"fmt"
)

//Create a new type of 'deck'
//which is a slice of stings
type deck []string

// d: The actual copy of the deck we're working that is available in the function
// In TypeScript it can be referenced as 'this' or in Python as 'self'
// any variable of type 'deck' now gets access to the 'print' method 
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}