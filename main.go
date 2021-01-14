package main

import (
	"fmt"
	"strings"
)

func main() {
	cards := newDeck()
	blackjack := 21

	n := 1
	for n < 10 {
		cards.shuffle()
		n++
	}

	hand := cards.dealTwo()
	hand.print()
	handValue := hand.value()
	fmt.Println("You're at ", handValue)

	fmt.Println("Hit or Stay?")
	var input string
	fmt.Scanln(&input)
	for strings.ToLower(input) == "hit" {
		newCard := cards.dealOne()
		hand = append(hand, newCard...)
		hand.print()
		handValue := hand.value()
		fmt.Println("You're at ", handValue)
		if handValue > 21 {
			fmt.Printf("You went over %v. You lose!", blackjack)
			break
		} else if handValue == blackjack {
			fmt.Printf("You got %v. You win!", blackjack)
			break
		}
		fmt.Println("Hit or Stay?")
		fmt.Scanln(&input)
	}
	handValue = hand.value()
	if handValue < blackjack {
		fmt.Println("You got", handValue)
		aiHand := play(cards)
		aiHandValue := aiHand.value()
		fmt.Println("The AI got", aiHandValue)
		if aiHandValue > handValue {
			fmt.Println("You loose!")
		} else if aiHandValue < handValue {
			fmt.Println("You win!")
		} else {
			fmt.Println("No one won...")
		}
	}
}
