package main

// import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()

}
