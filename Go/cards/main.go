package main

import (
	"fmt"
)

func main() {
	// var deckSize int = 50
	// fmt.Println(deckSize)

	// var cards []string
	// cards = append(cards, "x")
	// fmt.Println(cards)
	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("\n")
	remainingCards.print()

	// strCards := cards.toString()
	// fmt.Println(strCards)

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	// var c = []string{}
	// c = append(c, "s")
	// fmt.Println(c)
}

// NOTES:
// 1.) https://yourbasic.org/golang/for-loop/

// different declarations of slices
// x := []string{"1"}
// x = append(x, "2")
// for i, val := range x {
// 	fmt.Println(i, val)
// }
// fmt.Println(x)

// var c = []string{}
// c = append(c, "s")
// fmt.Println(c)
