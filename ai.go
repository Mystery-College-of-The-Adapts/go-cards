package main

import "fmt"

func play(d deck) deck {
	hand := d.dealTwo()
	handValue := hand.value()
	fmt.Println("The AI is at", handValue)
	for handValue < 16 {
		fmt.Println("The AI hit!")
		newCard := d.dealOne()
		hand = append(hand, newCard...)
		handValue = hand.value()
		fmt.Println("The AI is at", handValue)
	}
	fmt.Println("The AI stays.")
	return hand
}
