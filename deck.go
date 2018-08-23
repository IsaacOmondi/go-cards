package main

import (
	"time"
	"math/rand"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
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
// this function returns the number of cards in hand and the remaining cards 
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
// convert string slice to string
func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

// []byte() turns string into bytes
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
// Read File from local machine
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	// Convert byte to string and separate values with a comma
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// randomizes cards every time the program is run
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}