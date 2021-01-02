package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length to be 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spaces but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs but got %v", d[len(d)-1])
	}
}

func TestPrint(t *testing.T) {
	d := newDeck()
	d.print()
}

func TestDeal(t *testing.T) {
	d := newDeck()
	draw := 5
	hand, d := d.deal(draw)

	if len(hand) != draw {
		t.Errorf("Expected hand length to be 5 but got %v", len(hand))
	}

	if len(d) != (52 - draw) {
		t.Errorf("Expected deck length to be 47 but got %v", len(d))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	e := newDeck()
	e.shuffle()

	if d[0] == e[0] {
		t.Errorf("Expected first card to be random but got %v for both decks", d[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {

}
