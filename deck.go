package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamons", "Hearts", "Trebols"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
