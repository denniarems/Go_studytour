package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var cardValues = []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"  Of  "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
