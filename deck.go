package main

import (
	"fmt"
)

//Create a new type of 'deck'
//which is a slice of stings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// we use the underscore if the first variable in the for loop is not being used anywhere
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d: The actual copy of the deck we're working that is available in the function
// In TypeScript it can be referenced as 'this' or in Python as 'self'
// any variable of type 'deck' now gets access to the 'print' method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
