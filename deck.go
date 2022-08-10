package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	var cards deck
	cardSuits := []string{"s", "d", "h", "c"}
	cardValues := []string{"A", "2", "3", "4"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+suit)
		}
	}
	return cards
}

func (d deck) saveDeck(filename string) error {
	return ioutil.WriteFile(filename, []byte(strings.Join([]string(d), ",")), 0666)
}

func loadDeck(filename string) deck {
	arrBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error occured when trying to load a file")
		os.Exit(1)
	}
	fileText := string(arrBytes)
	return deck(strings.Split(fileText, ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	for i := range d {
		newPosition := rnd.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
