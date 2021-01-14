package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	symbol string
	suit   string
	value  int
}

type deck []card

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♠", "♥", "♦", "♣"}
	cardSymbolValues := map[string]int{
		"A":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  10,
		"Q":  10,
		"K":  10,
	}

	for _, suit := range cardSuits {
		for symbol, value := range cardSymbolValues {
			c := card{symbol, suit, value}
			cards = append(cards, c)
		}
	}

	return cards
}

func (d deck) print() {
	for _, c := range d {
		fmt.Println(c.symbol, c.suit)
	}
}

func (d *deck) dealOne() deck {
	*d = (*d)[1:]
	return (*d)[:1]
}

func (d *deck) dealTwo() deck {
	*d = (*d)[2:]
	return (*d)[:2]
}

func (d *deck) deal(handSize int) deck {
	*d = (*d)[handSize:]
	return (*d)[:handSize]
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}

func (d deck) value() int {
	v := 0
	for _, c := range d {
		v += c.value
	}
	return v
}

// func (d deck) toString() string {
// 	c := []card(d)
// 	return strings.Join([]string(c), ",")
// }

// func (d deck) saveToFile(filename string) error {
// 	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
// }

// func newDeckFromFile(filename string) deck {
// 	bs, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 		os.Exit(1)
// 	}
// 	s := strings.Split(string(bs), ",")
// 	return deck(s)
// }
