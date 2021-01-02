package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	for i, c := range cards {
		fmt.Println(i, c)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
