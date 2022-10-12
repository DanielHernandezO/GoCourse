package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diavaluemonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", " Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suite))
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%d %s \n", i, card)
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("There was an error getting the file")
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	for index := range d {
		newPosition := rand.Intn(len(d))
		fmt.Println(newPosition)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
