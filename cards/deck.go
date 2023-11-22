package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Diamonds", "Clubs", "Hearts", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack","Queen", "King"}
	
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

// Create a new function called 'print' which prints the deck
// This is a 'receiver' function for deck 
	// -> func (d deck) funcName() {...}
		// d is the actual copy of the deck we're working with
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
