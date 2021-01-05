package main

import (
	"fmt"
	"strings"
)

func main() {
	cards := newDeck()

	n := 1
	for n < 10 {
		cards.shuffle()
		n++
	}

	hand, cards := cards.dealTwo()
	hand.print()

	fmt.Println("Hit or Stay?")
	var input string
	fmt.Scanln(&input)
	for strings.ToLower(input) == "hit" {
		newCard, _ := cards.dealOne()
		hand = append(hand, newCard...)
		hand.print()
		handValue := hand.value()
		if handValue > 21 {
			fmt.Println("You went over 21. You lose!")
			break
		} else if handValue == 21 {
			fmt.Println("You got 21. You win!")
			break
		}
		fmt.Println("Hit or Stay?")
		fmt.Scanln(&input)
	}

}
