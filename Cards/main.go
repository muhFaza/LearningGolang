package main

import "fmt"

func main() {
	var newCards = newDeck()
	cardsDealt, hands, _ := newCards.dealCards(2)
	// fmt.Println(cardsDealt, newCards)
	// fmt.Println(newCards.dealCards(2))
	// cardsDealt.print()
	// newCards.print()
	fmt.Println(len(cardsDealt))
	fmt.Println(len(hands))
	fmt.Println(len(newCards))
	// fmt.Println([]byte(newCards[0]))
}