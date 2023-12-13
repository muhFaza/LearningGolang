package main

import (
	"errors"
	"fmt"
)

type deck []string

func (d deck) dealCards(n int) (deck, deck, error) {
	if n == 0 {
		return nil,nil, errors.New("No cards dealt")
	}

	maxIndex := len(d)
	cardsDealt := d[maxIndex-n :]
	cardsRemaining := d[:maxIndex-n]
	d = cardsRemaining

	return cardsDealt, cardsRemaining, nil
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cs := range cardSuits {
		for _, cv := range cardValues {
			cards = append(cards, (cv + " of " + cs))
		}
	}
	return cards
}

func (d deck) print() {
	for _, v := range d {
		fmt.Println(v)
	}
}
