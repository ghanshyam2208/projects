package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [...]string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := [...]string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func (d deck) printCards() {
	for index, card := range d {
		fmt.Println(index+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func deckFromFile(filename string) deck {
	return strings.Split(readFile(filename), ",")
}
