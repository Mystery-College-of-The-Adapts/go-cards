package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	l := len(d)
	if l != 52 {
		t.Errorf("Expected deck length to be 52 but got %v", l)
	}

	fc := d[0]
	if fc != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spaces but got %v", fc)
	}

	lc := d[len(d)-1]
	if lc != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs but got %v", lc)
	}
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

func TestToString(t *testing.T) {
	d := newDeck()
	ty := reflect.TypeOf(d.toString())
	if ty != reflect.TypeOf("") {
		t.Errorf("Expected deck to be of type string but got %v", ty)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	f := "_decktest"
	os.Remove(f)
	d := newDeck()
	d.saveToFile(f)
	nd := newDeckFromFile(f)

	l := len(nd)
	if l != 52 {
		t.Errorf("Expected deck length to be 52 but got %v", l)
	}

	os.Remove(f)

	nd = newDeckFromFile("_decktesting")
	if nd != nil {
		t.Error("Expected deck to be nil")
	}

}
