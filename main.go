package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	fmt.Println(cards)
	hand, cards := cards.dealOne()
	fmt.Println(hand)
}
