package main

import "fmt"

func main() {
	fmt.Println("Creating new deck of cards")
	deck := newDeck()
	fmt.Println("Printing deck of cards")
	deck.print()
	fmt.Println("Deal 3 cards from deck")
	hand, restOfDeck := deal(deck, 3)
	fmt.Println("Print the rest ff deck")
	restOfDeck.print()
	fmt.Println("Print hand")
	hand.print()
	fmt.Println("Shuffle cards in hand")
	hand.shuffle()
	fmt.Println("Print shuffled hand")
	hand.print()
}
