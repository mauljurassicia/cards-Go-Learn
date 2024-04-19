package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues  := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit )
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck, error){
  	if len(d) <= handSize {
		return nil, nil, errors.New("there are only " + strconv.Itoa(len(d)) + " cards")
	}
	return d[:handSize], d[handSize:], nil
}

func (d deck) toString() string{
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck{
	bs, err := os.ReadFile(fileName)
	if err != nil{
		fmt.Println("Error :", err.Error())
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle(){

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) -1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
